package main

import (
	"fmt"
	"net"
)

// RPC就是想实现函数调用模式的网络化。客户端就像调用本地函数一样，
// 然后客户端把这些参数打包之后通过网络传递到服务端，服务端解包到处理过程中执行，
// 然后执行的结果反馈给客户端。 RPC是基于Socket的
// RPC（Remote Procedure Call Protocol）——远程过程调用协议，是一种通过网络从远程计算机程序上请求服务，
// 而不需要了解底层网络技术的协议。它假定某些传输协议的存在，如TCP或UDP，以便为通信程序之间携带信息数据。
func RpcDemo() {

	/*
	Go标准包中已经提供了对RPC的支持，而且支持三个级别的RPC：TCP、HTTP、JSONRPC。
	但Go的RPC包是独一无二的RPC，它和传统的RPC系统不同，
	它只支持Go开发的服务器与客户端之间的交互，因为在内部，它们采用了Gob来编码。

	Go RPC的函数只有符合下面的条件才能被远程访问，不然会被忽略，详细的要求如下：
		函数必须是导出的(首字母大写)
		必须有两个导出类型的参数，
		第一个参数是接收的参数，第二个参数是返回给客户端的参数，第二个参数必须是指针类型的
		函数还要有一个返回值error
	正确的RPC函数格式如下：func (t *T) MethodName(argType T1, replyType *T2) error
		T、T1和T2类型必须能被encoding/gob包编解码。

	任何的RPC都需要通过网络来传递数据，Go RPC可以利用HTTP和TCP来传递数据，
	利用HTTP的好处是可以直接复用net/http里面的一些函数。
	具体看server/rpc_http_server.go和client/rpc_http_client.go文件，运行方式:
		1. 进入到目录cd D:\go_project\src\study_go\library\server，运行go run .\rpc_http_server.go
		2. 进入到目录cd D:\go_project\src\study_go\library\client，运行go run .\rpc_http_client.go
		3. 查看cmd的输出
	*/

	fmt.Println("具体看server/rpc_http_server.go和client/rpc_http_client.go文件")
}

// Socket起源于Unix，而Unix基本哲学之一就是“一切皆文件”，
// 都可以用“打开open –> 读写write/read –> 关闭close”模式来操作。
// Socket就是该模式的一个实现，网络的Socket数据传输是一种特殊的I/O，Socket也是一种文件描述符。
// 常用的Socket类型有两种：流式Socket（SOCK_STREAM）和数据报式Socket（SOCK_DGRAM）。
// 流式是一种面向连接的Socket，针对于面向连接的TCP服务应用；
// 数据报式Socket是一种无连接的Socket，对应于无连接的UDP服务应用。
// Socket是基于TCP和UDP协议的（TCP和UDP是基于IP协议的），
// 而Socket的上层是一些应用层调用，例如FTP，HTTP, RPC等都是基于Socket的上层应用协议
func SocketDemo() {
	addr := net.ParseIP("192.0.2.1")
	if addr == nil {
		fmt.Println("Invalid address")
	} else {
		fmt.Println("The address is ", addr.String())
	}

	// 具体参考chat包
}



/*  Web Socket简介

WebSocket是HTML5的重要特性，它实现了基于浏览器的远程socket，它使浏览器和服务器可以进行全双工通信

在WebSocket出现之前，为了实现即时通信，采用的技术都是“轮询”，即在特定的时间间隔内，
由浏览器对服务器发出HTTP Request，服务器在收到请求后，返回最新的数据给浏览器刷新，
“轮询”使得浏览器需要对服务器不断发出请求，这样会占用大量带宽。

WebSocket采用了一些特殊的报头，使得浏览器和服务器只需要做一个握手的动作，
就可以在浏览器和服务器之间建立一条连接通道。且此连接会保持在活动状态，
你可以使用JavaScript来向连接写入或从中接收数据，就像在使用一个常规的TCP Socket一样。
它解决了Web实时化的问题，相比传统HTTP有如下好处：
	一个Web客户端只建立一个TCP连接
	Websocket服务端可以推送(push)数据到web客户端.
	有更加轻量级的头，减少数据传送量

WebSocket的协议颇为简单，在第一次handshake通过以后，连接便建立成功，
其后的通讯数据都是以”\x00″开头，以”\xFF”结尾。WebSocket URL的起始输入是ws://或是wss://（在SSL上）。
在客户端，这个是透明的，WebSocket组件会自动将原始数据“掐头去尾”。

Go语言标准包里面没有提供对WebSocket的支持，但是在由官方维护的go.net子包中有对这个的支持，可以通过如下命令获取：
	go get code.google.com/p/go.net/websocket

*/


/* REST简介

REST(REpresentational State Transfer)这个概念，它指的是一组架构约束条件和原则。
满足这些约束条件和原则的应用程序或设计就是RESTful的。

Web应用要满足REST最重要的原则是: 客户端和服务器之间的交互在请求之间是无状态的,
即从客户端到服务器的每个请求都必须包含理解请求所必需的信息。
如果服务器在请求之间的任何时间点重启，客户端不会得到通知。
此外此请求可以由任何可用服务器回答，这十分适合云计算之类的环境。
因为是无状态的，所以客户端可以缓存数据以改进性能。

另一个重要的REST原则是系统分层，这表示组件无法了解除了与它直接交互的层次以外的组件。
通过将系统知识限制在单个层，可以限制整个系统的复杂性，从而促进了底层的独立性。

Go没有为REST提供直接支持，但是因为RESTful是基于HTTP协议实现的，
所以我们可以利用net/http包来自己实现，或者使用第三方的WEB框架来实现

*/