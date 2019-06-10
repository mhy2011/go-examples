package main

import "fmt"

func main() {
	ints := []int{1, 2, 3, 5, 6, 8}
	fmt.Println("slice =", ints)
	odd := filter(ints, isOdd)	//函数作为参数
	fmt.Println("odd =", odd)
	even := filter(ints, isEven)	//函数作为参数
	fmt.Println("even =", even)

}

// 声明一个函数类型
type TestInt func(int) bool

//判断一个整数是否为奇数
func isOdd(a int) bool {
	return !isEven(a)
}

// 判断一个整数是否为偶数
func isEven(a int) bool {
	if a%2 == 0 {
		return true
	}
	return false
}

// 声明的函数作为参数
func filter(slice []int, f TestInt) []int {
	var result []int
	for _, val := range slice {
		if f(val) {
			result = append(result, val)
		}
	}
	return result
}