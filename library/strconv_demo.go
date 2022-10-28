package main

import (
	"fmt"
	"strconv"
)

// 字符串转换
func StrconvDemo() {
	a, _ := strconv.ParseBool("false")
    b, _ := strconv.ParseFloat("123.23", 64)
    c, _ := strconv.ParseInt("1234", 10, 64)
    d, _ := strconv.ParseUint("12345", 10, 64)
    e, _ := strconv.Atoi("1023")
    fmt.Println(a, b, c, d, e)

	sa := strconv.FormatBool(false)
    sb := strconv.FormatFloat(123.23, 'g', 12, 64)
    sc := strconv.FormatInt(1234, 10)
    sd := strconv.FormatUint(12345, 10)
    se := strconv.Itoa(1023)
    fmt.Println(sa, sb, sc, sd, se)
}