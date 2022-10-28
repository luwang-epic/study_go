package main

import (
	"fmt"
	"net/rpc"
)

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

func main() {
	// 通过命令行获取参数，第一个参数为文件本身的地址，从第二个开始才是用户传入的
	// if len(os.Args) != 2 {
	// 	fmt.Println("Usage: ", os.Args[0], "server")
	// 	os.Exit(1)
	// }
	// serverAddress := os.Args[1]

	// 通过http协议来实现远程调用
	client, err := rpc.DialHTTP("tcp", "localhost:12345")
	if err != nil {
		fmt.Println("rpc dial http error:", err)
	}

	args := Args{17, 8}
	var reply int
	// 调用远程的方法，调用Arith对象的Multiply方法
	// 第1个要调用的函数的名字，第2个是要传递的参数，第3个要返回的参数(注意是指针类型)
	err = client.Call("Arith.Multiply", args, &reply)
	if err != nil {
		fmt.Println("arith invoke Multiply error:", err)
	}
	fmt.Printf("Arith Multiply: %d*%d=%d\n", args.A, args.B, reply)

	var quot Quotient
	err = client.Call("Arith.Divide", args, &quot)
	if err != nil {
		fmt.Println("arith invoke Divide error:", err)
	}
	fmt.Printf("Arith Divide: %d/%d=%d remainder=%d\n", args.A, args.B, quot.Quo, quot.Rem)
}