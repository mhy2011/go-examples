package main

import "fmt"

func main() {

	//定义匿名函数
	sumVal := func(x, y int) int {
		return x + y
	}
	//调用匿名函数
	i := sumVal(10, 20)
	fmt.Println("i =", i)

	// 定义并调用匿名函数
	mult := func(x, y int) int {
		return x * y
	}(10, 20)
	fmt.Println("mult =", mult)
}
