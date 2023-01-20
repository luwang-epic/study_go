package main

import (
	"fmt"
	"os/exec"
)

// 系统调用相关
func ExecDemo() {
	cmd := exec.Command("go", "version")
	output, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(output))
}

// 用于运行单元测试
func utSample(a int) int {
	return a
}

// 进入到当前目录，通过go run . 运行，这样会编译所有文件来运行
func main() {
	fmt.Println("演示go中常用库的使用")
	fmt.Println()

	// fmt库
	// FmtDemo()

	// strings库
	// StringsDemo()

	// 字符串转换相关
	// StrconvDemo()

	// 正则库regexp
	// RegexpDemo()

	// 时间库time
	// TimeDemo()

	// 数据计算相关的
	// MathDmoe()

	// 文件操作，大部分在os库下
	// OsDemo()

	// 日志相关
	LogDemo()
	zapLogDemo()

	// 系统调用相关
	// ExecDemo()

	// socket编程
	//SocketDemo()
}
