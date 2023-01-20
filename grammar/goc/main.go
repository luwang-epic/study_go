package main
//#include <stdio.h>
//void say1(){
// printf("Hello World\n");
//}
/*
#include <stdio.h>
void say2(){
	printf("Hello World\n");
}
*/
import "C"


/*
在Go语言开篇中我们已经知道, Go语言与C语言之间有着千丝万缕的关系, 甚至被称之为21世纪的C语言
所以在Go与C语言互操作方面，Go更是提供了强大的支持。尤其是在Go中使用C，
你甚至可以直接在Go源文件中编写C代码，这是其他语言所无法望其项背的
	在import "C"之前通过单行注释或者通过多行注释编写C语言代码
	在import "C"之后编写Go语言代码
	在Go语言代码中通过C.函数名称() 调用C语言代码即可
	注意: import "C"和前面的注释之间不能出现空行或其它内容, 必须紧紧相连

Go语言中没有包名是C的包, 但是这个导入会促使Go编译器利用cgo工具预处理文件
在预处理过程中,cgo会产生一个临时包, 这个包里包含了所有C函数和类型对应的Go语言声明
最终使得cgo工具可以通过一种特殊的方式来调用import "C"之前的C语言代码
*/

func main() {
	C.say1()
	C.say2()
}