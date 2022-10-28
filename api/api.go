package api

import "fmt"

func Hello() {
	fmt.Println("say hello...")
}

func Hi(name string) {
	fmt.Println("invoke Hi...")
	fmt.Println("say hi to", name)
}
