package main

import "fmt"

func main() {
	s := "abcd"
	fmt.Println("1. s address is", &s, ", value is", s)
	ptr := &s	//声明一个指针
	*ptr = "1234"	//修改指针所指向字符串的值
	fmt.Println("2. s address is", &s, ", value is", s)
	s = "ABCD"
	fmt.Println("3. s address is", &s, ", value is", s)
}
