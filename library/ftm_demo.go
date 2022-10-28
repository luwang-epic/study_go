package main

import "fmt"

// fmt库方法
func FmtDemo() {
	// Scan从标准输入扫描文本，读取由空白符分隔的值保存到传递给本函数的参数中，换行符视为空白符。
	// 本函数返回成功扫描的数据个数和遇到的任何错误。如果读取的数据个数比提供的参数少，会返回一个错误报告原因。

	// Scanf从标准输入扫描文本，根据format参数指定的格式去读取由空白符分隔的值保存到传递给本函数的参数中。
	// 本函数返回成功扫描的数据个数和遇到的任何错误。

	// Scanln类似Scan，它在遇到换行时才停止扫描。最后一个数据后面必须有换行或者到达结束位置。
	// 本函数返回成功扫描的数据个数和遇到的任何错误。

	// var str1, str2 string
	// num, err := fmt.Scanln(&str1, &str2)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(str1, str2, num)

	// var str3 string
	// num, err = fmt.Scan(&str3)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(str3, num)

	// print输出给定的字符串，每项之间自动加空格，如果是数值或字符，则输出对应的十进制表示
	// Println 自动在结尾输出\n,两个数值之间自动加空格，每项之间自动加空格
	// printf 格式化输出
	fmt.Print("a", "\n") // a
	fmt.Print('a', 'b', "\n") //输出97 98   字符之间会输出一个空格  
	fmt.Println("a")          //输出a  

	// 格式化相关的
	/*
	%
	v 		相应值的默认格式 				fmt.Printf("%v", name)		 {春生}
	%+v  	打印结构体时，会添加字段名 		fmt.Printf("%+v", people) 	main.Human{Name:"zhangsan"}
	%#v 	相应值的Go语法表示				fmt.Printf("%#v",people)	main.Human{Name:"春生"}
	%T 		相应值的类型的Go语法表示 		fmt.Printf("%T",people)		main.Human
	%%		字面上的百分号					fmt.Printf("%%")			%

	%t 		true 或 false					fmt.Printf("%t",true)		true

	%b		二进制表示
	%c		相应Unicode码点所表示的字符	
	%d		十进制表示
	%o		 八进制表示
	%x		 十六进制表示，字母形式为小写 a-f 
	%X		十六进制表示，字母形式为大写 A-F
	%e		科学计数法 e小写 (=%.6e) 6位小数点 
	%E		科学计数法 E大写
	%f 		(=%.6f) 6位小数点 有小数点而无指数
	%g		根据情况选择 %e 或 %f 以产生更紧凑的（无末尾的0）输出
	%G		根据情况选择 %E 或 %f 以产生更紧凑的（无末尾的0）输出

	%s	输出字符串表示（string类型或[]byte)
	%10s	输出字符串最小宽度为10(右对齐)
	%-10s	输出字符串最小宽度为10(左对齐)
	%5.10s	输出字符串最小宽度为5，最大宽度为10(右对齐)
	%-5.10s	输出字符串最小宽度为5，最大宽度为10(左对齐)
	%5.3s	输出字符串宽度为5,如果原字符串宽度大于3,则截断


	%p		地址十六进制表示，前缀 0x
	%#p		地址十六进制表示，不带前缀 0x
	*/
	fmt.Printf("%e\n", 10.2)
	fmt.Printf("%.3e\n", 10.2)
	fmt.Printf("%.3E\n", 10.2)
	fmt.Printf("%f\n", 10.2)

	fmt.Printf("%10s\n", "oldboy")
	fmt.Printf("%-10s\n", "oldboy")
	fmt.Printf("%.5s\n", "oldboy")
	fmt.Printf("%5.10s", "oldboy22222222")
}