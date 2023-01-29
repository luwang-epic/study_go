package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"os"
)

type Client struct {
	ServerIp   string
	ServerPort int
	Name string
	// 连接对象
	conn net.Conn

	// 当前客户端的模式，有：0：退出, 1：公聊模式, 2：私聊模式, 3：更新用户名
	flag int 
}

// 创建客户端
func NewClient(serverIp string, serverPort int) *Client {
	// 连接server
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", serverIp, serverPort))
	if err != nil {
		fmt.Println("net dial error:", err)
		return nil
	}

	client := &Client{
		ServerIp:   serverIp,
		ServerPort: serverPort,
		conn: conn,
		// 初始值不为0即可，否则直接退出了
		flag: 10,
	}
	return client
}

// 当前客户端支持的操作
func (client *Client) menu() bool {
	var flag int

	fmt.Println("1.公聊模式")
	fmt.Println("2.私聊模式")
	fmt.Println("3.更新用户名")
	fmt.Println("0.退出")

	// 从cmd读取输入，并赋给flag变量
	fmt.Scanln(&flag)

	if flag >= 0 && flag <= 3 {
		client.flag = flag
		return true
	}

	return false
}

// 运行客户端
func (client *Client) Run() {
	// 不为0一直运行
	for client.flag != 0 {
		// 如果返回的不是true，那么重新输入
		for client.menu() != true {
			fmt.Println(">>>>请输入合法范围内的数字<<<<<<<")
		}

		// 根据不同的模式处理不同的业务
		switch client.flag {
		case 1:
			fmt.Println("公聊模式选择...")
			client.PublicChat()
			break
		case 2:
			fmt.Println("私聊模式选择...")
			client.PrivateChat()
			break
		case 3:
			fmt.Println("更新用户名选择...")
			client.UpdateName()
			break
		}
	}
}

// 处理server回应的消息，直接显示到标准的输出即可
func (client *Client) DoResponse() {
	// 一旦client.conn有数据，就直接copy到stdout标准的输出上，永久阻塞监听
	io.Copy(os.Stdout, client.conn)

	// 和上面的类似
	// for {
	// 	buf := make([]byte, 4096)
	// 	client.conn.Read(buf)
	// 	fmt.Println(buf)
	// }
}

func (client *Client) UpdateName() bool {
	fmt.Println(">>>>请输入用户名:")
	fmt.Scanln(&client.Name)

	sendMsg := "rename|" + client.Name + "\n"
	_, err := client.conn.Write([]byte(sendMsg))
	if err != nil {
		fmt.Println("conn write err:", err)
		return false
	}

	return true
}

// 公聊实现
func (client *Client) PublicChat() {
	// 提示用户输入消息
	var chatMsg string
	fmt.Println(">>>>>请输入聊天内容, exit退出.")
	fmt.Scanln(&chatMsg)

	for chatMsg != "exit" {
		// 发送给服务器

		// 消息不为空则发送
		if len(chatMsg) !=0 {
			sendMsg := chatMsg + "\n"
			_, err := client.conn.Write([]byte(sendMsg))
			if err != nil {
				fmt.Println("conn write err:", err)
				break
			}
		}

		chatMsg = ""
		fmt.Println(">>>>>请输入聊天内容, exit退出.")
		fmt.Scanln(&chatMsg)
	}
}

// 私聊模式
func (client *Client) PrivateChat() {
	var remoteName string
	var chatMsg string

	client.FindOnlineUsers()
	fmt.Println(">>>>请输入聊天对象(用户名), exit退出...")
	fmt.Scanln(&remoteName)

	for remoteName != "exit" {
		fmt.Println(">>>>>请输入消息内容, exit退出:")
		fmt.Scanln(&chatMsg)

		for chatMsg != "exit" {
			// 消息不为空则发送
			if len(chatMsg) !=0 {
				sendMsg := "to|" + remoteName + "|" + chatMsg + "\n\n"
				_, err := client.conn.Write([]byte(sendMsg))
				if err != nil {
					fmt.Println("conn write err:", err)
					break
				}
			}

			chatMsg = ""
			fmt.Println(">>>>>请输入消息内容, exit退出:")
			fmt.Scanln(&chatMsg)
		}

		client.FindOnlineUsers()
		fmt.Println(">>>>请输入聊天对象(用户名), exit退出...")
		fmt.Scanln(&remoteName)
	}
}

// 查询在线的用户
func (client *Client) FindOnlineUsers() {
	sendMsg := "who\n"
	_, err := client.conn.Write([]byte(sendMsg))
	if err != nil {
		fmt.Println("conn write err:", err)
		return
	}
}

// 从执行语句中获取ip和port参数
var serverIp string
var serverPort int

// ./client.exe -ip 127.0.0.1 -port 8888
func init() {
	flag.StringVar(&serverIp, "ip", "127.0.0.1", "设置服务器IP地址(默认是127.0.0.1)")
	flag.IntVar(&serverPort, "port", 8888, "设置服务器IP地址(默认是8888)")
}

// 主函数
func main() {
	// 解析参数
	flag.Parse()
	fmt.Println("从命令行解析的服务器地址为:", serverIp, "端口为:", serverPort)

	client := NewClient(serverIp, serverPort)
	if client == nil {
		fmt.Println(">>>>>>> 连接服务器失败...")
		return
	}
	fmt.Println(">>>>>>> 连接服务器成功...")

	// 单独开启一个协程，处理server的回应消息
	go client.DoResponse()

	// 启动客户端业务
	client.Run()
}