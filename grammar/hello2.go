package main

import "fmt"

func main() {
	fmt.Println("study go")

	formatStr := "a=%d & b=%s"
	fmt.Println(fmt.Sprintf(formatStr, 1, "b"))

	fmt.Printf("变量的地址: %x\n", &formatStr)

	// 类型转换
	var sum int = 17
	var count int = 5
	var mean float32
	mean = float32(sum) / float32(count)
	fmt.Printf("mean 的值为: %f\n", mean)

	const (
		a = iota //0
		b        //1
		c        //2
		d = "ha" //独立值，iota += 1
		e        //"ha"   iota += 1
		f = 100  //iota +=1
		g        //100  iota +=1
		h = iota //7,恢复计数
		i        //8
	)
	const (
		j = 1 << iota
		k = 3 << iota
		m
		l
	)

	fmt.Println(a, b, c, d, e, f, g, h, i, j, k, m, l)

	// 闭包
	/* nextNumber 为一个函数，函数 i 为 0 */
	nextNumber := getSequence()

	/* 调用 nextNumber 函数，i 变量自增 1 并返回 */
	fmt.Println(nextNumber()) // 1
	fmt.Println(nextNumber()) // 2
	fmt.Println(nextNumber()) // 3

	/* 创建新的函数 nextNumber1，并查看结果 */
	nextNumber1 := getSequence()
	fmt.Println(nextNumber1()) // 1
	fmt.Println(nextNumber1()) // 2
}

func getSequence() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}
