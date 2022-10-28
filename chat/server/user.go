package main

import (
	"fmt"
	"net"
	"strings"
)

type User struct {
	Name string
	Addr string
	// 用户的消息缓冲区，从中读取消息，然后发送到客户端
	Channel chan string
	// 与服务器的连接
	conn net.Conn

	// 当前用户关联的server对象
	server *Server
}

// 创建一个用户
func NewUser(conn net.Conn, server *Server) *User {
	userAddr := conn.RemoteAddr().String()

	user := &User {
		Name: userAddr,
		Addr: userAddr,
		Channel: make(chan string),
		conn: conn,
		server: server,
	}

	// 启动监听
	go user.ListenMessage()

	return user
}

// 监听当前User channel，如果有消息，就直接发送给客户端
func (user *User) ListenMessage() {
	for {
		msg := <- user.Channel
		
		user.conn.Write([]byte(msg + "\n"))
	}
}


// 用户上线处理，广播上线消息
func (user *User) Online() {
	// 用户上线，将用户加入到server到map中
	user.server.mapLock.Lock()
	user.server.OnlineMap[user.Name] = user
	user.server.mapLock.Unlock()

	// 广播当前的用户消息
	user.server.BroadCast(user, "online")
}

// 用户下线处理，广播下线消息
func (user *User) Offline() {
	// 用户下线，将用户加入到server到map中
	user.server.mapLock.Lock()
	delete(user.server.OnlineMap, user.Name)
	user.server.mapLock.Unlock()

	// 广播当前的用户消息
	user.server.BroadCast(user, "offline")
}

// 处理用户发送来的消息
func (user *User) DoMessage(msg string) {
	// 如果输入的是who，就返回当前在线的用户有哪些
	if msg == "who" {
		fmt.Println("who...")
		user.server.mapLock.Lock()
		for _, cli := range user.server.OnlineMap {
			onlineMsg := "[" + cli.Addr + "] " + cli.Name + ": " + "online...\n"
			user.SendMsg(onlineMsg)
		}
		user.server.mapLock.Unlock()
	} else if len(msg) > 7 && msg[:7] == "rename|" {
		fmt.Println("user rename....")
		// 重命名用户
		// 消息格式rename|zhangsan
		newName := strings.Split(msg, "|")[1]

		// 判断用户是否存在
		_, ok := user.server.OnlineMap[newName]
		if ok {
			user.SendMsg("username already exist...")
		} else {
			user.server.mapLock.Lock()
			delete(user.server.OnlineMap, user.Name)
			user.server.OnlineMap[newName] = user
			user.server.mapLock.Unlock()

			renameMsg := "rename success from " + user.Name + " to " + newName + "...\n"
			user.Name = newName
			user.SendMsg(renameMsg)
		}
	} else if len(msg) > 4 && msg[:3] == "to|" {
		// 发送私聊消息，消息格式：to|zhangsan|content
		
		// 获取私聊的用户名
		remoteName := strings.Split(msg, "|")[1]
		if remoteName == "" {
			user.SendMsg("msg format is error, format is \"to|zhangsan|content\", please check")
			return
		}

		// 根据用户名，获取用户对象
		remoteUser, ok := user.server.OnlineMap[remoteName]
		if !ok {
			user.SendMsg("username isn't exist, please check...")
			return
		}

		//获取消息内容，并发送
		content := strings.Split(msg, "|")[2]
		if content == "" {
			user.SendMsg("invalid message without conent, please check")
			return
		}
		remoteUser.SendMsg(user.Name + " send message: " + content)
	} else {
		// 广播消息给每个用户
		user.server.BroadCast(user, msg)
	}
}

// 给当前的用户发送消息
func (user *User) SendMsg(msg string) {
	user.conn.Write([]byte(msg))
}
