package pkg2

import "fmt"

// 当前pkg2提供的函数
// 首字母大写，表示对外开放的函数，小写只能在当前包下调用
func Pkg2Api() {
	fmt.Println("pkg2 Pkg2Api() function execute ....")
}


func init() {
	fmt.Println("pkg2 init() function execute ....")
}