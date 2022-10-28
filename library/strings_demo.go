package main

import (
	"fmt"
	"strings"
	"unicode"
)

// strings包来进行搜索(Contains、Index)、替换(Replace)和解析(Split、Join)等操作，
// 但是这些都是简单的字符串操作，他们的搜索都是大小写敏感，而且固定的字符串，
// 如果我们需要匹配可变的那种就没办法实现了，当然如果strings包能解决你的问题，那么就尽量使用它来解决。
// 因为他们足够简单、而且性能和可读性都会比正则好。
func StringsDemo() {
	fmt.Println("************************strings包使用*****************************")
	
	// 字符串查找相关

	s := "hello word"
	// 返回字符串s包含字符串substr的个数
	fmt.Printf("字符串:%s,word出现数量: %d\n", s, strings.Count(s, "word"))
	// 判断字符串s是否包含substr字符串
	fmt.Printf("字符串:%s 是否包含word: %t \n", s, strings.Contains(s, "word"))
	fmt.Printf("字符串:%s 是否包go: %t \n", s, strings.Contains(s, "go"))
	// 判断字符串s是否包含chars字符串中的任意一个字符
	fmt.Printf("字符串:%s 是否包含go中的任意一个字符: %t \n", s, strings.ContainsAny(s, "go"))
	fmt.Printf("字符串:%s 是否包含gg中的任意一个字符: %t \n", s, strings.ContainsAny(s, "gg"))

	// 返回字符串s中字符串substr最后一次出现的位置
	fmt.Printf("在字符串%s中,字符串o最后一次出现的位置: %d \n", s, strings.LastIndex(s, "o"))
	// 返回字符串s中字符串substr首次出现的位置
	fmt.Printf("在字符串%s中,字符串o首次出现的位置: %d \n", s, strings.Index(s, "o"))
	// 返回字符串s中字符c首次出现的位置
	var b byte = 'l'
	fmt.Printf("在字符串%s中,字符l首次出现的位置: %d \n",s, strings.IndexByte(s, b))
	// 返回字符串s中字符c最后一次出现的位置
	fmt.Printf("在字符串%s中,字符l最后一次出现的位置: %d \n", s, strings.LastIndexByte(s, b))

	// 判断字符串s是否有前缀prefix
	a := "VIP001"
	fmt.Printf("字符串:%s 是否有前缀vip: %t \n", a, strings.HasPrefix(a, "vip"))
	fmt.Printf("字符串:%s 是否有前缀VIP: %t \n", a, strings.HasPrefix(a, "VIP"))

	// 判断字符串s是否有后缀suffix
	sn := "K011_Mn"
	fmt.Printf("字符串:%s 是否有后缀MN: %t \n", sn, strings.HasSuffix(sn, "MN"))
	fmt.Printf("字符串:%s 是否有后缀Mn: %t \n", sn, strings.HasSuffix(sn, "Mn"))
	// 返回字符串s中满足函数f(r)==true,字符首次出现的位置 (判断第一个汉字的位置)
	f := func(c rune) bool {
	return unicode.Is(unicode.Han,c)
	}
	s4 := "go!中国人"
	fmt.Printf("字符串:%s 首次出现汉字的位置: %d \n", s4, strings.IndexFunc(s4,f))
	fmt.Printf("字符串:%s 最后一次出现汉字的位置: %d \n", s4, strings.LastIndexFunc(s4,f))


	//  字符串分割相关

	s5 := "Go! Go! 中国人!"
	// 将字符串s以空白字符分割，返回切片
	slice := strings.Fields(s5)
	fmt.Printf("将字符串:%s以空白字符分割, 返回切片:%v \n", s5, slice)
	// 将字符串s以满足f(r)==true的字符分割，分割后返回切片。
	// 以特殊符号分割
	ff := func(r rune) bool {
		// 不是字母，也不是数字
		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	}
	ss := "张三@19*BeiJing&高中生|男(打球"
	slice2 := strings.FieldsFunc(ss, ff)

	fmt.Printf("将字符串:%s 以满足f(r)==true的字符分割, 返回切片:%v \n", ss, slice2)

	// 将字符串s以sep作为分割符进行分割，分割后字符最后去掉sep
	s2 := "@123@张@AB@001"
	sep := "@"
	slic1 := strings.Split(s2, sep)
	fmt.Printf("将字符串:【%s】以%s进行分割, 分割后最后去掉:%s 返回切片: %v 切片长度: %d \n",s2, sep, sep, slic1,len(slic1))

	// 将字符串s以sep作为分割符进行分割，分割后字符最后加上sep,返回切片
	slic2 := strings.SplitAfter(s2, sep)
	fmt.Printf("将字符串:【%s】以%s进行分割, 分割后最后加上:%s 返回切片: %v 切片长度: %d \n",s2, sep, sep, slic2,len(slic2))
	// 将字符串s以sep作为分割符进行分割，分割后字符最后加上sep,n决定分割成切片长度
	fmt.Printf("将字符串:【%s】以%s进行分割, 指定分割切片长度%d: %v 分割后加上%s \n",s2, sep, 0, strings. SplitAfterN(s2, sep, 0), sep)
	fmt.Printf("将字符串:【%s】以%s进行分割, 指定分割切片长度%d: %v 分割后加上%s \n",s2, sep, 1, strings. SplitAfterN(s2, sep, 1), sep)
	// 将字符串s以sep作为分割符进行分割，分割后字符最后去掉sep, n决定分割成切片长度
	fmt.Printf("将字符串:【%s】以%s进行分割, 指定分割切片长度%d: %v 分割后去掉%s \n",s2, sep, 3, strings.SplitN(s2, sep, 3), sep)
	
	// 字符串大小写相关

	str := "hello word"
	str1 := "HELLO WORD"
	// Title(s string) string: 每个单词首字母大写
	// fmt.Printf("Title->将字符串%s 每个单词首字母大写: %s\n", str, strings.Title(str))
	// ToLower(s string) string : 将字符串s转换成小写返回
	fmt.Printf("ToLower->将字符串%s 转换成小写返回: %s\n",str1,strings.ToLower(str1))
	// ToTitle(s string)string: 将字符串s转换成大写返回
	// 大部分情况下， ToUpper 与 ToTitle 返回值相同，但在处理某些unicode编码字符则不同
	fmt.Printf("ToTitle->将字符串%s 转换成大写返回: %s\n",str,strings.ToTitle(str))
	// ToUpper(s string)string: 将字符串s转换成大写返回
	fmt.Printf("ToUpper->将字符串%s 转换成大写返回: %s\n",str,strings.ToUpper(str))

	// 字符串拼接相关

	// 字符串拼接
	fmt.Printf("字符串拼接:Join-> %s\n",strings.Join([]string{"a","b","c"},"|"))
	// 字符串重复多少次
	fmt.Printf("字符串重复:Repeat-> %s\n",strings.Repeat("Go!", 10))

	//  字符串替换

	// 字符串替换,如果n<0会替换所有old子串。
	s = "a,b,c,d,e,f"
	old := ","
	newStr := "."
	fmt.Printf("将字符串【%s】中的前%d个【%s】替换为【%s】结果是【%s】\n", s, 2, old, newStr,strings.Replace(s, old, newStr, 2))
	fmt.Printf("将字符串【%s】中的前%d个【%s】替换为【%s】结果是【%s】\n", s, 7, old, newStr,strings.Replace(s, old, newStr, 7))
	fmt.Printf("将字符串【%s】中的前%d个【%s】替换为【%s】结果是【%s】\n", s, -1, old, newStr,strings.Replace(s, old, newStr, -1))
	// 字符串全部替换
	fmt.Printf("将字符串【%s】中的【%s】全部替换为【%s】结果是【%s】\n", s, old, newStr, strings.ReplaceAll(s, old, newStr))

	// 字符串比较

	// 字符串比较大小
	s = "a"
	s1 := "c"
	s2 = "c"
	fmt.Printf("%s < %s 返回 : %d \n", s, s1, strings.Compare(s, s1))
	fmt.Printf("%s > %s 返回 : %d \n", s1, s, strings.Compare(s1, s))
	fmt.Printf("%s = %s 返回 : %d \n", s1, s2, strings.Compare(s1, s2))
	// 字符串比较一致性
	sa := "go"
	sb := "Go"
	sc := "go"
	fmt.Printf("%s和%s是否相等(忽略大小写): %t \n", sa, sb, strings.EqualFold(sa, sb))
	fmt.Printf("%s和%s是否相等(忽略大小写): %t \n", sa, sc, strings.EqualFold(sa, sc))
	fmt.Printf("%s和%s是否相等(不忽略大小写): %t \n", sa, sb, sa == sb)
	fmt.Printf("%s和%s是否相等(不忽略大小写): %t \n", sa, sc, sa == sc)

	// 字符串删除

	// 将字符串首尾包含在cutset中的任一字符去掉
	str = "@*test@-@124@!*"
	cutset := "*#@!"
	fmt.Printf("将字符串【%s】首尾包含在【%s】中的任一字符去掉,返回:【%s】\n",str, cutset, strings.Trim(str, cutset))
	// 将字符串首尾满足函数`f(r)==true`的字符串去掉
	fff := func(r rune) bool {
		return strings.Contains("*#@!", string(r))
	}
	fmt.Printf("将字符串【%s】首尾满足函数f的字符去掉,返回:【%s】\n", str, strings.TrimFunc(str, fff))
	// 将字符串左边包含在cutset中的任一字符去掉
	fmt.Printf("将字符串【%s】左边包含在【%s】中的任一字符去掉,返回:【%s】\n",str, cutset, strings.TrimLeft(str, cutset))
	// 将字符串左边满足函数`f(r)==true`的字符串去掉
	fmt.Printf("将字符串【%s】左边满足函数f的字符去掉,返回:【%s】\n", str, strings.TrimLeftFunc(str, fff))
	// 将字符串右边包含在cutset中的任一字符去掉
	fmt.Printf("将字符串【%s】右边包含在【%s】中的任一字符去掉,返回:【%s】\n", str, cutset, strings.TrimRight(str, cutset))
	fmt.Printf("将字符串【%s】右边满足函数f的字符去掉,返回:【%s】\n", str, strings.TrimRightFunc(str, fff))
   
	// 将字符串中前缀字符串prefix去掉
	str1 = "VIP00001_U"
	fmt.Printf("将字符串【%s】前缀【%s】去掉,返回:【%s】\n", str1, "VIP", strings.TrimPrefix(str1,"VIP"))
	fmt.Printf("将字符串【%s】前缀【%s】去掉,返回:【%s】\n", str1, "vip", strings.TrimPrefix(str1,"vip"))
   
	 // 将字符串中后缀字符串suffix去掉
	fmt.Printf("将字符串【%s】后缀【%s】去掉,返回:【%s】\n", str1, "U", strings.TrimSuffix(str1,"U"))
	fmt.Printf("将字符串【%s】后缀【%s】去掉,返回:【%s】\n", str1, "u", strings.TrimSuffix(str1,"u"))
   
	// 将字符串首尾空白去掉
	str2 := "  hello  word !  "
	fmt.Printf("将字符串【%s】首尾空白去掉,返回:【%s】\n", str2, strings.TrimSpace(str2))

	fmt.Println()
	fmt.Println()
}