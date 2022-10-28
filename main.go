package main

import (
	"fmt"
	"github.com/luwang-epic/study_go/api"
	"github.com/luwang-epic/study_go/api/external"
	// "github.com/luwang-epic/study_go/api/internal"
)

// 演示如何访问项目内的包
func viewProjectPackage() {
	// 访问对外api包下的函数
	api.Hi("lisi")

	// 访问对外api包子包下面的函数
	api2.ExternalApi()

	// 是有的包访问不到，编译报错
	// use of internal package github.com/luwang-epic/study_go/api/internal not allowed
	// api3.InternalApi()
}

func main() {
	fmt.Println("study go")

	viewProjectPackage()
}
