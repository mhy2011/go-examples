package main

import "fmt"

func main() {
	a, b := 10, 20
	i := sum(a, b)	//调用函数
	fmt.Printf("%d + %d = %d\n", a, b, i)
}

// 定义一个函数
func sum(a, b int) int {
	return a + b
}