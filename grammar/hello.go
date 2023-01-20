// 表示该文件所属的包名
// 包名称，不需要有相关的文件夹的
package main

// 表示引入一个包，导包会依此执行该包的导入操作，以及全局的一些常量和变量，以及init()方法等
import "fmt"

/*
也可以用这种方式导包，导入自定义的包
*/
import (
	"github.com/luwang-epic/study_go/grammar/package/pkg1"
	"github.com/luwang-epic/study_go/grammar/package/pkg2"
)

// 如果想导入某个包，但是不想使用包中的方法，那么使用匿名的导入，会执行init方法
import _ "github.com/luwang-epic/study_go/grammar/package/pkg3"

// 也可以给导入包取一个别名
import pkg "github.com/luwang-epic/study_go/grammar/package/pkg4"

// 也可以通过.的方式导入包，这样可以直接使用导入的包函数
// 尽量不使用这种方式来进行导包
import . "github.com/luwang-epic/study_go/grammar/package/pkg5"

import (
	"encoding/json"
	"errors"
	"reflect"
	"runtime"
	"sync"
	"time"
)

/*
Go语言编码风格
1.go程序编写在.go为后缀的文件中
2.包名一般使用文件所在文件夹的名称; 包名应该简洁、清晰且全小写
3.main函数只能编写在main包中
4.每一条语句后面可以不用编写分号(推荐)
5.如果没有编写分号,一行只能编写一条语句
6.函数的左括号必须和函数名在同一行
7.导入包但没有使用包编译会报错
8.定义局部变量但没有使用变量编译也会报错
9.定义函数但没有使用函数不会报错
10.给方法、变量添加说明,尽量使用单行注释
*/

// 单个包中代码执行顺序: main包-->常量-->全局变量-->init函数-->main函数-->Exit

// main是函数名，是一个主函数和程序的入口
// 函数的左花括号，一定要和函数名在同一行，否则编译出错
func main() {
	// 官方更加推荐使用单行注释,而非多行注释，可以参考一些源码函数的注释

	// 调用函数输出
	// go中的表达式可以加分号，也可以不加，建议不加
	fmt.Println("hello, world")

	// 定义变量
	varAndConstDeclaration()

	// go的所有关键字
	keywordDemo()

	// go的类型
	typeDemo()

	// 声明函数
	funcDemo()

	// 使用其他包的函数
	pkg1.Pkg1Api()
	pkg2.Pkg2Api()

	// 通过包别名调用
	pkg.Pkg4Api()
	// 不需要包名
	Pkg5Api()

	// 指针
	point()

	// defer关键字使用
	deferDemo()

	// 数组和动态数组
	arrayDemo()

	// map类型
	mapDemo()

	// 异常处理
	exceptionDemo()

	// 结构体类型
	structDemo()

	// 接口
	interfaceDemo()

	// reflect反射
	reflectDemo()

	// 锁
	lockDemo()

	// 协程机制，m:n模型，m个协程对应n个内核线程；需要有调度器来调度m个协程到n个内核线程中
	// goroutine协程，会有多个调度器（可以配置GOMAXPROCES），每个调度器有一个队列，负责调度队列中的协程执行
	// 也有一个全局的协程队列，防止需要执行的协程
	// 特点：复用线程，多个并行，抢占，全局G队列
	routineDemo()

}

/*********************协程**************************/
func routineDemo() {
	// 创建一个go协程，去执行newTask()流程
	// 如果main退出，那么这个协程也会退出
	go newTask()

	// 用go创建承载一个形参为空，返回值为空的一个函数
	go func() {
		defer fmt.Println("A.defer")

		// 定义一个参数为s的匿名函数，返回值为bool类型
		b := func(s string) bool {
			defer fmt.Println("B.defer", s)
			return true
		}("bbb") // 调用匿名函数

		fmt.Println("匿名函数的返回值:", b)
		if b {
			// 退出当前gorouine
			runtime.Goexit()
		}
		fmt.Println("BBB")
	}()

	// go协程之间通信，通过管道通信
	// 定义一个channel，int类型，无缓存
	c := make(chan int)
	go func() {
		defer fmt.Println("func gorotine 结束")
		fmt.Println("func goroutine 正在运行....")

		// 将数据放入到管道中
		// 如果管道没有缓冲，那么会阻塞；
		// 如果管道有缓冲，没有达到容量，不会阻塞，如果达到容量，会阻塞
		c <- 666
	}()

	// 从管道中取数据，数据还没有过来，会阻塞；数据过来后会唤醒
	num := <-c
	fmt.Println("num =", num)

	// 带有缓冲的channel，3为缓冲的容量
	cc := make(chan int, 3)
	fmt.Println("缓冲的大小:", len(cc), "缓存的容量:", cap(cc))

	go func() {
		for i := 0; i < 5; i++ {
			cc <- i
		}

		// 关闭一个channel
		// 关闭channel后是可以读取数据，直到读取不到数据
		close(cc)
	}()
	for {
		// ok如果为true表示channel没有关闭，如果为false，表示channel已经关闭
		if data, ok := <-cc; ok {
			fmt.Println("从有缓冲的channel中获取数据:", data)
		} else {
			// 跳出循环
			break
		}
	}

	fc := make(chan int)
	quit := make(chan int)

	// 子线程中获取数据
	go func() {
		for i := 0; i < 10; i++ {
			// 如果fc中没有数据，将会是阻塞状态，等待写入数据
			fmt.Println("获取fibonacii的值为:", <-fc)
		}

		// 退出
		quit <- 0
	}()
	// 计算fibonacci的值，并放入到fc中，上面从fc中获取，都是阻塞的形式
	fibonacci(fc, quit)

	ccc := make(chan int, 3)
	go func() {
		for i := 0; i < 5; i++ {
			ccc <- i
		}
		close(ccc)
	}()
	// 可以使用range不断的从channel中读取数据，没有数据时，会阻塞在这里
	for data := range ccc {
		fmt.Println("不断的从channel中读取数据:", data)
	}
	fmt.Println("range读取数据结果")

	// 死循环
	for {
		time.Sleep(10 * time.Minute)
	}
}

func newTask() {
	i := 0
	for {
		i++
		fmt.Println("newTask goroutine: i =", i)
		time.Sleep(1 * time.Second)
		if i == 10 {
			// 退出协程
			runtime.Goexit()
		}
	}
}

func fibonacci(c, quit chan int) {
	fmt.Println("fibonacii...")
	x, y := 1, 1
	for {
		// 循环的判断
		select {
		case c <- x:
			// 如果c可写，将会进入这个case
			// 当channel中的数据取走时，将是可写状态
			y = x + y
			x = y - x
		case <-quit:
			// 如果quit可读，进入这个case，退出循环
			fmt.Println("quit...")
			return
		}
	}
}

/**********************锁***************************/

func lockDemo() {
	var lock sync.Mutex

	// 可以不进行显示的初始化，会自动初始化
	// lock = sync.Mutex{}

	lock.Lock()
	fmt.Printf("lock, %T\n", lock)
	fmt.Println(lock)
	lock.Unlock()

	rwlock := sync.RWMutex{}
	rwlock.Lock()
	fmt.Printf("rwlock, %T\n", rwlock)
	fmt.Println(rwlock)
	rwlock.Unlock()

	rwlock.RLock()
	fmt.Println(rwlock)
	rwlock.RUnlock()
}

/*********************反射**************************/

func reflectDemo() {
	var f float64 = 1.2345
	fmt.Println("类型是：", reflect.TypeOf(f), "值为：", reflect.ValueOf(f))

	user := User{1, "zhangsan", 22}
	GetFieldAndInvokeMethod(user)

	// 标签在json中的应用
	movie1 := Movie{"喜剧之王", 2000, 1.123, []string{"zhangsan", "lisi"}}
	jsonStr, err := json.Marshal(movie1)
	if err != nil {
		fmt.Println("json marshal error", err)
		return
	}
	fmt.Println(reflect.TypeOf(jsonStr))
	fmt.Printf("jsonStr: %s\n", jsonStr)

	movie2 := Movie{}
	err = json.Unmarshal(jsonStr, &movie2)
	if err != nil {
		fmt.Println("json unmarshal error", err)
		return
	}
	fmt.Println(movie2)
}

func GetFieldAndInvokeMethod(input interface{}) {
	// 获取input中的type
	inputType := reflect.TypeOf(input)
	fmt.Println("input type is", inputType)

	// 获取input中的value
	inputValue := reflect.ValueOf(input)
	fmt.Println("input value is", inputValue)

	// 通过type获取里面的字段
	// 1. 获取interface的reflect.Type, 通过Type得到NumField，进行遍历
	// 2. 得到每个field数据类型
	// 3. 通过field有一个Interface方法得到对应的value
	for i := 0; i < inputType.NumField(); i++ {
		field := inputType.Field(i)
		value := inputValue.Field(i).Interface()
		fmt.Println("field name:", field.Name, ", field type:", field.Type, ", field value:", value)

		// 获取字段中的标签名称
		tagInfo := field.Tag.Get("info")
		tagDoc := field.Tag.Get("doc")
		fmt.Println("info:", tagInfo, ", doc:", tagDoc)
	}

	// 通过type获取你们的方法，并调用
	for i := 0; i < inputType.NumMethod(); i++ {
		m := inputType.Method(i)
		fmt.Println("method name:", m.Name, "method type:", m.Type)

	}
}

type User struct {
	// 标签的典型应用是json和orm框架
	// 给Id定义info和doc标签，并给出具体的值，标签名称和值中间不能有空格
	Id   int    `info:"name" doc:"我的名字"`
	Name string `info:"sex"`
	Age  int
}

func (user User) Call() {
	fmt.Println("call...", user)
}

// 字段和json中的名称对应
type Movie struct {
	Title  string   `json:"title"`
	Year   int      `json:"year"`
	Price  float64  `json:"price"`
	Actors []string `json:"actors"`
}

/*********************接口**************************/
func interfaceDemo() {
	var animal Animal
	animal = &Cat{"green"}
	// 调用的就是Cat的Sleep方法，多态的现象
	animal.Sleep()

	animal = &Dog{"yellow"}
	// 调用的就是Dog的Sleep方法，多态的现象
	animal.Sleep()

	cat := Cat{"green"}
	dog := Dog{"yellow"}
	showAnimal(&cat)
	showAnimal(&dog)

	// 可以用interface{}来引用任何类型的数据
	baseInterface(10)
	baseInterface("abc")

	// 每个对象内部都维护了一个Pair<type, value>的结构体
	// f: pair<type:File, value:aaa>
	f := File{"aaa"}

	// r: pair<type:空, value:空>
	var r Reader
	// r: pair<type:File, value: aaa>
	r = f
	r.ReadFile(f.filename)

	// w: pair<type:空, value:空>
	var w Writer
	// w: pair<type:File, value: aaa>
	w = r.(Writer) // 此处w r具有的type类型一直，都是File对象，因为File对象实现了Reader和Writer接口
	w.WriteFile(f.filename)
}

func baseInterface(arg interface{}) {
	fmt.Println("baseInterface is called with ", arg)

	// interface{}中提供了 类型断言 的机制
	value, ok := arg.(string) // 判断是否是string类型，只有interface{}可以这么用
	if ok {
		fmt.Println("arg is string type, value =", value)
	}

	// Go里面switch默认相当于每个case最后带有break
	// 匹配成功后不会自动向下执行其他case，而是跳出整个switch
	switch arg.(type) {
	case string:
		fmt.Println("arg is string type")
	case int:
		fmt.Println("arg is int type")
	default:
		fmt.Println("arg is not int or string type")
	}

	return
}

func showAnimal(animal Animal) {
	fmt.Println("animal type is ", animal.GetType())
}

// 定义动物接口
type Animal interface {
	Sleep()
	GetColor() string
	GetType() string
}

// go中只需要实现接口的所有方法，就表示该类实现了这个接口
// 定义Cat来实现Animal接口
type Cat struct {
	color string
}

func (cat *Cat) Sleep() {
	fmt.Println("Cat Sleep...")
}
func (cat *Cat) GetColor() string {
	return cat.color
}
func (cat *Cat) GetType() string {
	return "Cat"
}

// 定义Dog来实现Animal接口
type Dog struct {
	color string
}

func (dog *Dog) Sleep() {
	fmt.Println("Dog Sleep...")
}
func (dog *Dog) GetColor() string {
	return dog.color
}
func (dog *Dog) GetType() string {
	return "Dog"
}

type File struct {
	filename string
}

func (file File) ReadFile(filename string) {
	fmt.Println("read file...", file.filename)
}
func (file File) WriteFile(filename string) {
	fmt.Println("write file...", file.filename)
}

type Reader interface {
	ReadFile(filename string)
}
type Writer interface {
	WriteFile(filename string)
}

/*********************结构体**************************/

func structDemo() {
	var a myint = 10
	fmt.Printf("a = %d, type of a = %T\n", a, a)

	var book Book
	book.title = "t"
	book.auth = "a"
	structFunc1(book)
	fmt.Println(book)

	structFunc2(&book)
	fmt.Println(book)

	hero := Hero{Name: "h", Ad: 10, level: 1}
	fmt.Println(hero.Name, hero.Ad, hero.level)
	hero.SetName1("hh")
	fmt.Println(hero)
	hero.SetName2("hhhh")
	fmt.Println(hero)

	man := SuperMan{Man{"man", 11}, 33}
	man.walk()
	man.eat()

}

// 声明一种数据类型myint，是int的一个别名
type myint int

// 定义一个结构体
type Book struct {
	title string
	auth  string
}

// 传递的是一个Book的副本
func structFunc1(book Book) {
	book.auth = "bbbb"
}

// 传递的是一个Book的引用
func structFunc2(book *Book) {
	book.auth = "cccc"
}

// 如果类名首字母大写，表示其他包也可以访问
type Hero struct {
	// 如果属性的首字母大写，表示属性是对外能够访问的，否则的话只能够本包访问
	Name  string
	Ad    int
	level int
}

// 这个方法是Hero的方法, this是一个名称，可以自己决定；首字母大写表示外部可以访问的
func (this Hero) Show() {
	fmt.Println(this)
}
func (hero Hero) GetName() string {
	return hero.Name
}
func (this Hero) SetName1(name string) {
	//this是调用这个方法对象的一个副本（拷贝）
	this.Name = name
}
func (this *Hero) SetName2(name string) {
	//this是调用这个方法对象的引用
	this.Name = name
}

type Man struct {
	name string
	age  int
}

func (man Man) walk() {
	fmt.Println("Man walk")
}
func (man Man) eat() {
	fmt.Println("Man eat")
}

type SuperMan struct {
	Man // 继承子类

	level int
}

// 重写子类的方法
func (superMan *SuperMan) eat() {
	fmt.Println("SuperMan eat")
}

/********************异常处理***********************/
// go不推荐使用panic和recover耐久进行异常处理，特别是对外的方法中
// 因为调用者并不知道你是否抛出了方法，是否需要捕获，会造成混乱
// go推荐让函数返回error的方式来抛出异常，error是go中的一个接口
// 可以自己实现error接口来定义自己的异常，也可以通过errors包来使用定义好的异常

func exceptionDemo() {
	recoverDemo()

	normalReturn, err := errorDemo()
	fmt.Println("正常返回:", normalReturn, "异常信息:", err)
	// 一般这么使用
	if err != nil {
		fmt.Println("出现错误")
		return
	}
	// 否则正常执行
}

func recoverDemo() {
	defer fmt.Println("这个defer也是可以执行的")

	// 必须要先声明defer，否则不能捕获到panic异常
	defer func() {
		// 使用recover捕获异常
		err := recover()
		if err != nil {
			fmt.Println("捕获异常:", err)
		}
		fmt.Println("异常捕获后...可以执行到这里")
	}()

	defer fmt.Println("这个defer是可以执行的")

	// 异常部分需要放到defer之后
	panicDemo()

	fmt.Println("异常之后的代码，执行不到这里")
}

func panicDemo() {
	a := 10
	if a > 2 {
		// 使用panic抛出异常
		panic("出现异常，抛出去")
	}
}

func errorDemo() (int, error) {
	err := errors.New("this is error")
	return 1, err
}

/********************map类型**************************/
func mapDemo() {
	var map1 map[string]string
	if map1 == nil {
		fmt.Println("map1是一个空的map")
	}

	// 在使用map前，需要下先用make给map分配数组空间
	map1 = make(map[string]string, 10)
	map1["a"] = "java"
	map1["b"] = "go"
	fmt.Println(map1)

	// 在容量不够是也会动态扩容
	map2 := make(map[int]string)
	map2[0] = "java"
	map2[1] = "go"
	// map不能使用cap函数
	fmt.Println("长度为：", len(map2), "内容为：", map2)

	map3 := map[string]int{
		"one": 1,
		"two": 2,
	}
	map3["three"] = 3
	delete(map3, "one")
	map3["two"] = 22
	for key, value := range map3 {
		fmt.Println(key, value)
	}

	mapFunc(map3)
	fmt.Println(map3)
}

// map是一个引用传递
func mapFunc(map1 map[string]int) {
	map1["four"] = 4
}

/********************数组和动态数组**************************/
func arrayDemo() {
	// 固定长度的数组
	var arr1 [10]int
	for i := 0; i < 10; i++ {
		fmt.Print(arr1[i], " ")
		arr1[i] = 10
		fmt.Print(arr1[i], " ")
	}
	fmt.Println()

	arr2 := [4]int{1, 2, 3}
	for index, value := range arr2 {
		fmt.Println("index =", index, "value = ", value)
	}

	// 动态数组，动态数组的本身就是指向后面数组的指针，本身就是一个引用传递
	// 在go中叫切片
	arr3 := []int{11, 22}

	// 固定长度的数组，类型为[n]int，因此作为函数的参数时候需要加上长度，不同长度的数组不能互相传递，
	// 固定长度的数组也是值传递
	fmt.Printf("type of arr1 = %T\n", arr1)
	fmt.Printf("type of arr2 = %T\n", arr2)
	fmt.Printf("type of arr3 = %T\n", arr3)
	arrayFunc(arr1, arr3)
	// 还是10，并没有改为100
	fmt.Println("arr1[0] = ", arr1[0])
	// _表示匿名的变量，不需要使用下标时，可以如下循环
	for _, value := range arr3 {
		fmt.Println("value = ", value)
	}

	//  声明slice1是一个切片，长度为3，并且初始化为1,2,3
	slice1 := []int{1, 2, 3}
	fmt.Println("len of slice1 = ", len(slice1))

	// 声明slice2是一个切片，但是没有给slice分配空间
	var slice2 []int
	fmt.Println("len of slice2 = ", len(slice2))

	// 声明slice3是一个切片，同时给slice分片3个空间，初始值都是0
	var slice3 []int = make([]int, 3)
	fmt.Println("len of slice3 = ", len(slice3))

	// 声明slice4是一个切片，同时给slice4分片3个空间，初始值是0，通过:=推导出slice是一个切片
	slice4 := make([]int, 3)
	fmt.Println("len of slice4 = ", len(slice4))

	// 判断一个slice是否有空间
	if slice2 == nil {
		fmt.Println("slice2是没有空间的")
	}

	// 申请一个长度为3，容量为5的数组，长度表示左右指针之间的距离，容量表示左指针到底层数组末尾之间的距离
	slice5 := make([]int, 3, 5)
	fmt.Println("长度为：", len(slice5), "容量为：", cap(slice5), "内容为：", slice5)
	slice5 = append(slice5, 1)
	fmt.Println("长度为：", len(slice5), "容量为：", cap(slice5), "内容为：", slice5)
	slice5 = append(slice5, 2)
	fmt.Println("长度为：", len(slice5), "容量为：", cap(slice5), "内容为：", slice5)
	// 容量不够时，将会将容量增加为2倍
	slice5 = append(slice5, 3)
	fmt.Println("长度为：", len(slice5), "容量为：", cap(slice5), "内容为：", slice5)

	// 底层还是对应相同的地址空间，因此改变其中之一，另一个值也会改变
	s1 := slice5[0:3] // [0, 3)
	fmt.Println(s1)
	s1[0] = 10
	fmt.Println(s1)
	fmt.Println(slice5)

	s2 := make([]int, 2)
	// 复制数组，将s1中的值赋值到s2
	copy(s2, s1)
	s2[0] = 1000
	fmt.Println(s2)
	fmt.Println(s1)
}

// arr1固定长度数组是值传递，而不是引用传递
// arr3动态数组是引用传递
func arrayFunc(arr1 [10]int, arr3 []int) {
	arr1[0] = 100
	arr3[0] = 100
}

/***************defer使用*************************/

func deferDemo() int {
	// defer使用，在函数执行完成后再调用这个语句，会采用压栈的方式，先定义的后执行
	defer fmt.Println("main end1")
	defer fmt.Println("main end2")

	defer deferFunc()
	return returnFunc()
}

func deferFunc() int {
	fmt.Println("deferFunc called....")
	return 0
}

func returnFunc() int {
	fmt.Println("returnFunc called....")
	return 0
}

/****************************指针************************************/
func point() {
	var a, b int = 1, 2
	// 将a的地址传给函数
	swap(&a, &b)
	fmt.Println("通过指针改变值: a = ", a, "b = ", b)

	var p *int
	p = &a
	fmt.Println(&a)
	fmt.Println(p)

	// 二级指针
	var pp **int
	pp = &p
	fmt.Println(&p)
	fmt.Println(pp)
}

func swap(pa, pb *int) {
	// 改变p所对应的地址的值
	var temp int = *pa
	*pa = *pb
	*pb = temp
}

/*************************函数的声明*********************************/

func funcDemo() {
	r1 := oneReturnValue(10)
	fmt.Println(r1)

	r1, r2 := twoReturnValue1(1, "a")
	fmt.Println(r1, r2)

	r1, r2 = twoReturnValue2(1, "a")
	fmt.Println(r1, r2)
}

// 返回一个返回值
func oneReturnValue(a int) int {
	return 10 + a
}

// 返回多个值，匿名的
func twoReturnValue1(a int, b string) (int, int) {
	fmt.Println(a, b)
	return 66, 77
}

// 返回多个值，有形参名称的
func twoReturnValue2(a int, b string) (r1 int, r2 int) {
	// r1,r2属于这个函数的形参，会进行初始化为0值，作用域为该函数
	fmt.Println(r1, r2)

	// 给有名称的返回值赋值
	r1 = 101
	r2 = 102

	// 这个会将r1，r2的值返回
	return
}

// 如果类型相同，可以这样写
func twoReturnValue3(a int, b string) (r1, r2 int) {
	// 给有名称的返回值赋值
	r1 = 101
	r2 = 102

	// 这个会将r1，r2的值返回
	return
}

/****************************go语言的类型**********************************/
/*
go语言的类型
基本类型
	数值类型
		整型
			有符号
				int (默认, 和CPU平台的字长一样)
				int8
				int16
				int32
				int64
				rune (表示Unicode字符对应的整型)
			无符号
				uint (和CPU平台的字长一样))
				uint8
				uint16
				uint32
				uint64
				uintptr
					它可以保存一个指针地址，它可以进行指针运算，
					是一个能足够容纳指针大小的整数类型。
					一般用于底层和c交互的编程
		浮点型
			单精度型 float32
			双精度型 float64
		复数型
			complex64
				实数部分和虚数部分都是 float32 类型
			complex128
				实数部分和虚数部分都是 float64 类型
	字符类型 byte 是uint8的别名，表示一个字节  定义是这样的：type byte = uint8
	布尔类型 bool
		true
		false
	字符串类型 string

派生类型
	指针类型(Pointer)  *a, 取指针用&a
	数组类型 [2]int 不同的长度表示不同的类型
	切片类型 []int 也就是动态数组
		切片是一个拥有相同类型元素的可变长度的序列。它是基于数组类型做的一层封装。
		它非常灵活，支持自动扩容。切片是一个引用类型，它的内部结构包含地址、长度和容量。
	Map类型 map[string]string
	结构体(struct)
	管道类型 chan string
	函数类型 函数本身也是一种数据类型, 函数类型的作用是实现函数的多态  func(int32, int32) int32
	接口类型(interface)


数据类型转换格式: 数据类型(需要转换的数据)
*/

func typeDemo() {
	var c uint8 = 'a'
	// 97
	fmt.Println(c)
	// a
	fmt.Printf("%c", c)

	// a
	fmt.Println(string(97))
	// 你
	fmt.Println(string(20320))
	temp := []rune{20320, 22909, 32, 19990, 30028}
	// 你好 世界
	fmt.Println(string(temp))

	// 对于英文字符串，不管是用rune类型还是byte类型，不管是字符串的长度还是取值，都是相同的
	// 对于中文字符来说，rune类型的操作就比byte类型的操作更加友好很多，
	// 	我们可以通过[:]操作直接取出中文的对应数量，而byte取出来却是乱码??
	s := "hello world"
	// [104 101 108 108 111 32 119 111 114 108 100]
	fmt.Println("byte=", []byte(s))
	// [104 101 108 108 111 32 119 111 114 108 100]
	fmt.Println("rune=", []rune(s))
	// he
	fmt.Println(s[:2])
	// he
	fmt.Println(string([]rune(s))[:2])
	ss := "你好 世界"
	// [228 189 160 229 165 189 32 228 184 150 231 149 140]
	fmt.Println("byte=", []byte(ss))
	// [20320 22909 32 19990 30028]
	fmt.Println("rune=", []rune(ss))
	// 乱码
	fmt.Println(ss[:2])
	// 你好
	fmt.Println(ss[:6])
	// 你好
	fmt.Println(string([]rune(ss)[:2]))

	// 复数
	var score complex64 = complex(1, 2)
	var number complex128 = complex(23.23, 11.11)
	fmt.Println("Real Score = ", real(score), " Image Score = ", imag(score))
	fmt.Println("Real Number = ", real(number), " Image Number = ", imag(number))

	f := add1
	r := f(1, 2)
	fmt.Println("函数调用的结果为:", r)
	fmt.Printf("变量的类型为: %T\n", f)
	f = add2
	r = f(1, 2)
	fmt.Println("函数调用的结果为:", r)

	var ff func(int, int) int = add1
	fmt.Println("函数调用的结果为:", ff(1, 2))
}

func add1(a int, b int) int {
	return a + b
}

func add2(a int, b int) int {
	return a + a + b + b
}

/****************************go语言的关键字说明**********************************/

/*
关键字：
package：包
import：导入包
var：定义变量
const：不可修改常量
type：定义类型
struct：定义结构体
interface：定义接口
func：定义函数
return：返回
select：go语言特有的channel选择结构
chan：定义go语言的channel
defer：栈空间结束时候执行，类似于析构
go：开始并发执行
map：map类型
range：从slice、map等结构中取元素
if：选择结构
else：选择
goto：跳转
switch：选择
case：选标签择
default：当分支没有选择好的时候，默认使用的分支
fallthrough：fallthrough会不管下一层的case条件直接进行执行
for：循环
break：跳出循环
continue：跳过下面语句，继续执行循环

内建常量:
true	false	iota	nil

內建函数
make	申请空间 make([]int, 2)
len		获取长度 len(arr)
cap		获取容量 cap(arr)
new		new只分配内存，make用于slice，map，和channel的初始化  new(int)
append	添加元素 append(slice, 1)
copy	把一个切片内容复制到另一个切片中 copy(目标切片, 源切片)
delete	删除元素 delete(map, 'a')
real	获取复数的实部
imag	获取复数的虚部
panic	用来抛出异常
recover	用来捕获异常，从而恢复正常代码执行，recover必须配合defer使用
complex	定义复数 complex(1,2)：实部为1，虚部为2的复数
*/

func keywordDemo() {
	var n = 20
	if n > 10 {
		// Go语言的goto语句可以无条件地转移到程序中指定的行。
		// goto语句通常与条件语句配合使用。可用来实现条件转移，跳出循环体等功能。
		// 在Go程序设计中一般不主张使用goto语句， 以免造成程序流程的混乱，使理解和调试程序
		goto lable1
	}
	fmt.Println("不走这里")
lable1:
	fmt.Println("走这里")

lable2:
	for n > 10 {
		// 不改变值，会一直循环在这里，不会像java的while true一样使用goto跳出循环后就不执行循环了
		n = 5
		goto lable2
	}
	fmt.Println("结束...", n)

	/*
		Go里面switch默认相当于每个case最后带有break，匹配成功后不会自动向下执行其他case，而是跳出整个switch,
		但是可以使用fallthrough强制执行后面的case代码。

		switch中的break和fallthrough语句的区别：
			break：可以使用在switch中，也可以使用在for循环中
				强制结束case语句，从而结束switch分支
			fallthrough：用于穿透switch
				当switch中某个case匹配成功之后，就执行该case语句
				如果遇到fallthrough，那么后面紧邻的case，无需匹配， 执行穿透执行。
				fallthrough应该位于某个case的最后一行;如果它出现在中间的某个地方，编译器就会抛出错误。
	*/

	s := 1
	switch s {
	case 1:
		fmt.Println("执行case1")
	case 2:
		fmt.Println("不执行case2")
	}

	switch s {
	case 1:
		fmt.Println("执行case1")
		// break之后，后面的代码不会执行
		break
		fmt.Println("再次执行case1")
	case 2:
		fmt.Println("不执行case2")
	}

	switch s {
	case 1:
		fmt.Println("执行case1")
		fallthrough
	case 2:
		fmt.Println("执行case2")
	case 3:
		fmt.Println("不执行case3")
	}
}

/*******************变量和常量的声明方式*****************************/

// 全局变量，全局变量不能用:=的方式声明
var g int = -1

// 定义全局的常量，只读，不能修改
const cg int = 10

func varAndConstDeclaration() {
	// 声明方式一：使用默认初始值，为0
	var a int
	// 声明方式二：赋予初始值
	var b int = 10
	// 声明方式三：不加类型；不推荐，推荐使用方式二
	var c = 100
	// 声明方式四：使用:=的方式；这种方式只支持在函数中使用，不能在全局中使用
	// 在函数中声明推荐使用这种方式
	d := 1000
	d = 1001
	d = d + a
	fmt.Printf("a = %d; b = %d; c = %d; d = %d\n", a, b, c, d)

	// 声明变量方式五：一次声明多个变量
	var e, f int = 101, 102
	var ee, ff = 103, "aaa"
	fmt.Printf("e = %d, f= %d, ee = %d, ff = %s", e, f, ee, ff)

	// 声明变量方式六：使用多行
	var (
		s  string = "abc"
		ss string = "aabbcc"
		bb bool   = true
	)
	fmt.Printf("s = %s, ss = %s, bb = %s, type of cc = %T, type of bb = %T\n", s, ss, s, bb, bb)

	// 定义常量
	const con int = 10

	const (
		// 可以在const()添加一个关键字iota，每行的iota都会累加1， 第一行的iota的默认值是0
		// iota只能配合const的多行一起使用，实现累加效果，不能用于其他地方
		BEIJING  = iota // iota = 0
		SHANGHAI        // iota = 1
		SHENZHEN        // iota = 2
	)
	fmt.Printf("BEIJING = %d, SHANGHAI = %d, SHENMZHEN = %d\n", BEIJING, SHANGHAI, SHENZHEN)

	const (
		ca, cb = iota + 1, iota + 2 // iota = 0, ca = iota + 1 = 1, cb = iota + 2 = 2
		cc, cd                      // iota = 1, cc = iota + 1 = 2, cd = iota + 2 = 2
		ce, cf                      // iota = 2, ce = iota + 1 = 3, cf = iota + 2 = 4

		cg, ch = iota * 2, iota * 3 // iota = 3, cg = iota * 2 = 6, ch = iota * 3 = 9
		ci, ck                      // iota = 4, ci = iota * 2 = 8, ck = iota * 3 = 12
	)
	fmt.Println("ca = ", ca, "cb = ", cb, "cc = ", cc, "cd = ", cd, ce, cf, cg, ch, ci, ck)
}
