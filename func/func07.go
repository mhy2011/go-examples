package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "abc"
	fmt.Printf("s address is:%v, value is:%s\n", &s, s)
	Upper(s)
	fmt.Printf("after call Upper s address is:%v, value is:%s\n", &s, s)
}

/**
字符串是值拷贝，对字符串的变更不会影响到原定义的字符串
 */
func Upper(str string) (res string) {
	fmt.Printf("str address is:%v, value is:%s\n", &str, str)
	str = strings.ToUpper(str)
	res = str
	fmt.Printf("after to upper str address is:%v, value is:%s\n", &str, str)
	return
}
