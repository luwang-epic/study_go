package main

import (
	"errors"
	"fmt"
	"net/http"
	"net/rpc"
)

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

// 远程调用对象
type Arith struct {

}

// 也可以用下面的基本类型来作为对象
// type Arith int

func (t *Arith) Multiply(args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

func (t *Arith) Divide(args *Args, quo *Quotient) error {
	if args.B == 0 {
		return errors.New("divide by zero")
	}
	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B
	return nil
}

func main() {
	fmt.Println("rpc server by http starting...")

	arith := new(Arith)
	// 注册了一个Arith的RPC服务
	rpc.Register(arith)
	// 把该服务注册到了HTTP协议上，然后我们就可以利用http的方式来传递数据了。
	rpc.HandleHTTP()

	// 启动服务
	err := http.ListenAndServe(":12345", nil)
	if err != nil {
		fmt.Println("http ListenAndServe error", err)
	}
}