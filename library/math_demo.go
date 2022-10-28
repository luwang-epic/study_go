package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func MathDmoe() {
	// 自然底数
	fmt.Println(math.E)
	// 圆周率
	fmt.Println(math.Pi)
	fmt.Println(math.MaxInt)
	fmt.Println(math.MinInt)
	fmt.Println(uint64(math.MaxUint64))

	fmt.Println("是否是NaN:", math.IsNaN(math.NaN()), math.NaN())
	fmt.Println("是否是Inf:", math.IsInf(math.Inf(0), 0), math.Inf(0))
	fmt.Println("是否是Inf:", math.IsInf(math.Inf(-1), -1), math.Inf(-1))

	fmt.Println(math.Ceil(2.5))   // 3
	fmt.Println(math.Floor(2.5))  // 2
	fmt.Println(math.Ceil(-2.5))  // -2
	fmt.Println(math.Floor(-2.5)) // -3
	// 取正数部分
	fmt.Println(math.Trunc(2.5))  // 2
	fmt.Println(math.Trunc(-2.5)) // -2
	// 取正数和小数部分
	fmt.Println(math.Modf(2.5))  // 2 0.5
	fmt.Println(math.Modf(-2.5)) // -2 -0.5
	fmt.Println(math.Abs(-2.5))  // 2.5
	fmt.Println(math.Max(2, 1))  // 2
	fmt.Println(math.Min(2, 1))  // 1
	fmt.Println(math.Dim(3, 7))  // 0  x - y > 0 ? x -y : 0
	fmt.Println(math.Dim(7, 3))  // 4

	fmt.Println(math.Mod(2, 5))   // 2
	fmt.Println(math.Sqrt(4))     // 2
	fmt.Println(math.Pow(3, 2))   // 9
	fmt.Println(math.Pow10(2))    // 100
	fmt.Println(math.Exp(2))      // = pow(math.E, 2)
	fmt.Println(math.Log(math.E)) // 1
	fmt.Println(math.Log1p(4))    // = log(4+1)
	fmt.Println(math.Log(4 + 1))  // 1.6094379124341003
	fmt.Println(math.Log2(4))     // 2
	fmt.Println(math.Log10(100))  // 2

	fmt.Println(math.Sin(1))
	fmt.Println(math.Cos(1))
	fmt.Println(math.Tan(1))

	// 全局的随机数
	rand.Seed(time.Now().UnixMilli())
	fmt.Println(rand.Int())
	fmt.Println(rand.Intn(10))
	fmt.Println(rand.Int31n(20))
	arr := []int{1, 2, 3, 4}
	// 打乱数组
	rand.Shuffle(len(arr), func(i, j int) {
		arr[i], arr[j] = arr[j], arr[i]
	})
	fmt.Println(arr)

	// 也可以自己生成一个随机数生成器
	selfSource := rand.NewSource(time.Now().UnixMilli())
	selfRand := rand.New(selfSource)
	fmt.Println(selfRand.Intn(10))

}