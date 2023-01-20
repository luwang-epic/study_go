package pkg1

import "fmt"

// 当前pkg1提供的函数
func Pkg1Api() {
	fmt.Println("pkg1 Pkg1Api() function execute ....")
}


func init() {
	fmt.Println("pkg1 init() function execute ....")
}