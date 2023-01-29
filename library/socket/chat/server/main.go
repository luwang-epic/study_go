package main

import "fmt"

func main() {
	fmt.Println("一个简单的聊天系统")

	// 启动服务
	server := NewServer("127.0.0.1", 8888)
	server.Start()

	/*
	可以通过go build -o server.exe main.go server.go user.go编译
		然后直接通过./server.exe运行程序
	也可以通过go run .\main.go .\server.go .\user.go来运行服务端

	客户端有两种：
		window的cmd客户端
			可以通过如下方式启动cmd客户端：注意cmd的目录为D:\java\local\netcat-win32-1.12
				D:\java\local\netcat-win32-1.12> .\nc.exe 127.0.0.1 8888
		自己写的client.go的客户端
			可以通过go build -o client.exe client.go编译客户端
				然后执行运行client.exe
			或者也可以直接go run client.go来运行客户端
	*/
}