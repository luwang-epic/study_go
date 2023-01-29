package main

import (
	"fmt"
	"io"
	"net"
	"sync"
	"time"
)

// 服务器类，提供启动和处理聊天的能力
type Server struct {
	Ip string
	Port int

	// 在线用户列表
	OnlineMap map[string]*User
	// 锁
	mapLock sync.RWMutex

	// 消息广播的管道
	Message chan string
}

// 创建一个服务器
func NewServer(ip string, port int) *Server {
	server := &Server{
		Ip: ip,
		Port: port,
		OnlineMap: make(map[string]*User),
		Message: make(chan string),
	}

	return server
}

// 启动服务器
func (server *Server) Start() {
	// socket listen
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", server.Ip, server.Port))
	if err != nil {
		fmt.Println("net listen err:", err)
		return
	}
	// close socket
	defer listener.Close()

	// 启动监听服务器
	go server.ListenMessage()

	// 等待客户端连接，并处理请求
	for {
		// accept, 阻塞等待
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener accept err:", err)
			continue
		}

		// 在协程中处理业务逻辑，主线程中继续等待客户端连接
		go server.Hanlder(conn)
	}
}

func (server *Server) Hanlder(conn net.Conn) {
	// 处理具体的逻辑
	fmt.Println("链接建立成功....", conn.RemoteAddr())

	// 用户上线
	user := NewUser(conn, server)
	user.Online()

	// 判断用户是否活跃的channel
	isLive := make(chan bool)

	// 接受用户发送过来的消息
	go func() {
		buf := make([]byte, 4096)
		for {
			n, err := conn.Read(buf)
			if n == 0 {
				fmt.Println("关闭链接成功....", user.Addr)
				user.Offline()
				return
			}

			if err != nil && err != io.EOF {
				fmt.Println("conn read err:", err)
				return
			}

			// 提取用户的消息（去除最后的\n字符，因此取[0,n-1)，不包括第n-1个元素）
			msg := string(buf[:n-1])

			// 用户处理得到的消息
			user.DoMessage(msg)

			// 用户的任意消息，代表当前用户是一个活跃的用户
			isLive <- true
		}
	}()

	// 当前handler阻塞
	for {
		select {
		case <- isLive:
			// 当前用户是活跃的，应该重置定时器
			// 不做任何事，为了激活select，更新下面的定时器

		case <- time.After(time.Minute * 5):
			// 已经超时，将当前的用户强制关闭
			user.SendMsg("you already timeout, close...")

			// 将用户从server中移除
			// 接受用户消息的协程中已经做了，当读取不到消息时会将用户从server中移除

			// 销毁资源
			close(user.Channel)

			// 关闭连接
			conn.Close()

			// 退出当前的handler
			return
		}
	}
	
}

// 广播消息
func (server *Server) BroadCast(user *User, msg string) {
	sendMsg := "[" + user.Addr + "] " + user.Name + ": " + msg

	// 发送消息到缓冲区
	server.Message <- sendMsg
}

// 监听Message广播消息channel的gorouine，一旦有消息就发送给全部的在线User
func (server *Server) ListenMessage() {
	for {
		// 从channel中取消息，没有就阻塞
		msg := <- server.Message

		// 将msg发送给全部的在线用户
		server.mapLock.Lock()
		for _, user := range server.OnlineMap {
			user.Channel <- msg
		}
		server.mapLock.Unlock()
	}
}

