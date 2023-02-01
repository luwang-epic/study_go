package main

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"time"
)

/*
context可以用来在goroutine之间传递上下文信息，相同的context可以传递给运行在不同goroutine中的函数，
上下文对于多个goroutine同时使用是安全的，context包定义了上下文类型，可以使用background、TODO创建一个上下文，
在函数调用链之间传播context，也可以使用WithDeadline、WithTimeout、WithCancel 或 WithValue 创建的修改副本替换它，
听起来有点绕，其实总结起就是一句话：context的作用就是在不同的goroutine之间同步请求特定的数据、取消信号以及处理请求的截止日期。

目前我们常用的一些库都是支持context的，例如gin、database/sql等库都是支持context的，
这样更方便我们做并发控制了，只要在服务器入口创建一个context上下文，不断透传下去即可。

context包主要提供了两种方式创建context:
	context.Backgroud()
	context.TODO()
这两个函数其实只是互为别名，没有差别，官方给的定义是：
	context.Background 是上下文的默认值，所有其他的上下文都应该从它衍生（Derived）出来。
	context.TODO 应该只在不确定应该使用哪种上下文时使用；
所以在大多数情况下，我们都使用context.Background作为起始的上下文向下传递。

上面的两种方式是创建根context，不具备任何功能，具体实践还是要依靠context包提供的With系列函数来进行派生：
	func WithCancel(parent Context) (ctx Context, cancel CancelFunc)
	func WithDeadline(parent Context, deadline time.Time) (Context, CancelFunc)
	func WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc)
	func WithValue(parent Context, key, val interface{}) Context
*/

func ContextDemo() {
	uuid := uuid.New().String()
	fmt.Println("uuid -> ", uuid)
	ctx1 := context.WithValue(context.Background(), "trace_id", uuid)
	go printLogDemo(ctx1)
	time.Sleep(1 * time.Second)

	ctx2, cancel2 := context.WithTimeout(context.Background(), 3 * time.Second)
	defer cancel2()
	timeoutDemo(ctx2, cancel2)

	ctx3, cancel3 := context.WithCancel(context.Background())
	go cancelDemo(ctx3)
	time.Sleep(10*time.Second)
	cancel3()
	time.Sleep(1*time.Second)
}


// 我们日常在业务开发中都希望能有一个trace_id能串联所有的日志，这就需要我们打印日志时能够获取到这个trace_id，
// 在python中我们可以用gevent.local来传递，在java中我们可以用ThreadLocal来传递，
// 在Go语言中我们就可以使用Context来传递，通过使用WithValue来创建一个携带trace_id的context，
// 然后不断透传下去，打印日志时输出即可
func printLogDemo(ctx context.Context) {
	// 在使用withVaule时要注意四个事项：
	// 1. 不建议使用context值传递关键参数，关键参数应该显示的声明出来，不应该隐式处理，context中最好是携带签名、trace_id这类值。
	// 2. 因为携带value也是key、value的形式，为了避免context因多个包同时使用context而带来冲突，key建议采用内置类型。
	// 3. 上面的例子我们获取trace_id是直接从当前ctx获取的，实际我们也可以获取父context中的value，在获取键值对是，我们先从当前context中查找，没有找到会在从父context中查找该键对应的值直到在某个父context中返回 nil 或者查找到对应的值。
	// 4. context传递的数据中key、value都是interface类型，这种类型编译期无法确定类型，所以不是很安全，所以在类型断言时别忘了保证程序的健壮性。
	traceId, ok := ctx.Value("trace_id").(string)
	if ok {
		fmt.Printf("trace_id = %s\n", traceId)
	} else {
		fmt.Println("get trace_id failure...")
	}
}

// 通常健壮的程序都是要设置超时时间的，避免因为服务端长时间响应消耗资源，
// 所以一些web框架或rpc框架都会采用withTimeout或者withDeadline来做超时控制，
// 当一次请求到达我们设置的超时时间，就会及时取消，不在往下执行。withTimeout和withDeadline作用是一样的，
// 就是传递的时间参数不同而已，他们都会通过传入的时间来自动取消Context，这里要注意的是他们都会返回一个cancelFunc方法，
// 通过调用这个方法可以达到提前进行取消，不过在使用的过程还是建议在自动取消后也调用cancelFunc去停止定时减少不必要的资源浪费。
//
// withTimeout、WithDeadline不同在于WithTimeout将持续时间作为参数输入而不是时间对象，这两个方法使用哪个都是一样的，
// 看业务场景和个人习惯了，因为本质withTimout内部也是调用的WithDeadline。
func timeoutDemo(ctx context.Context, cancel context.CancelFunc) {
	for i:=0; i< 10; i++ {
		time.Sleep(1*time.Second)
		select {
		case <- ctx.Done():
			fmt.Println("超时了...", ctx.Err())
			return
		default:
			fmt.Printf("do something %d\n", i)
			// 如果没有到时间就想结束，可以调用cancel取消
			//cancel()
			//fmt.Println("手动取消继续执行...")
		}
	}

	/*
	既可以超时自动取消，又可以手动控制取消。这里大家要记的一个坑，就是我们往从请求入口透传的调用链路中的context是携带超时时间的，
	如果我们想在其中单独开一个goroutine去处理其他的事情并且不会随着请求结束后而被取消的话，
	那么传递的context要基于context.Background或者context.TODO重新衍生一个传递，否决就会和预期不符合了
	 */
}

// 日常业务开发中我们往往为了完成一个复杂的需求会开多个gouroutine去做一些事情，这就导致我们会在一次请求中开了多个goroutine确无法控制他们，
// 这时我们就可以使用withCancel来衍生一个context传递到不同的goroutine中，当我想让这些goroutine停止运行，就可以调用cancel来进行取消。
func cancelDemo(ctx context.Context) {
	for range time.Tick(time.Second){
		select {
		case <- ctx.Done():
			fmt.Println("我要闭嘴了...")
			return
		default:
			fmt.Println("说话...balabalabalabala")
		}
	}
}
