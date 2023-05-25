package pkg3

import "fmt"

func init() {
	fmt.Println("pkg3 init() function execute ....")
}

// 可以定义多个init()方法
func init() {
	fmt.Println("pkg3 another init() function execute ....")
}

// Init 大写的Init不会被执行
func Init() {
	fmt.Println("pkg3 Init() function execute ....")
}


func Pkg3Api() {
	fmt.Println("pkg3 Pkg3Api() function execute ....")
}