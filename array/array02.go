package main

import "fmt"

const SIZE int = 5

func main() {
	var ary [SIZE]string
	fmt.Printf("1. ary address is %p\n", &ary)
	// 声明数组进指定了数组大小，则在调用的函数在声明形参时也需要指定数组大小，要么都声明，要么都不声明
	array2Upper(ary)
}

func array2Upper(ary [SIZE]string)  {
	fmt.Printf("2. ary address is %p\n", &ary)
}

