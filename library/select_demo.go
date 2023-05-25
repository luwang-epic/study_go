package main

import (
	"fmt"
)

var concurrencyCh chan struct{}

func SelectDemo() {
	fmt.Println("select demo...")

	concurrencyCh = make(chan struct{}, 5)
	concurrentLimitDemo(10)

	// 如果不放在select中，这里会阻塞
	//concurrencyCh <- struct{}{}

}

func concurrentLimitDemo(limit int) {
	for i := 0; i < limit; i++ {
		select {
		// 这里不会阻塞，如果没有满了，会选择defualt执行
		case concurrencyCh <- struct{}{}:
			fmt.Println("do something...")
		default:
			fmt.Println("concurrent limit...")
		}
	}
}