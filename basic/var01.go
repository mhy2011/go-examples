package main

import "fmt"

func main() {
	//初始化一个变量a
	var a uint8 = 20
	//变量a的值赋值为变量b, 实际上是内存值拷贝
	var b = a
	fmt.Printf("address for a is:%v, value is:%d\n", &a, a)
	fmt.Printf("address for b is:%v, value is:%d\n", &b, b)
}
