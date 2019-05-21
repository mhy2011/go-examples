package main

import "fmt"

func main() {
	// 使用make创建切片, len=cap
	numbers := make([]int, 5)
	fmt.Printf("numbers.address=%p, len=%d, cap=%d, val=%v\n",
		numbers, len(numbers), cap(numbers), numbers)
	// 使用make创建切片, len!=cap
	numbers2 := make([]int, 5, 8)
	fmt.Printf("numbers2.address=%p, len=%d, cap=%d, val=%v\n",
		numbers2, len(numbers2), cap(numbers2), numbers2)
}
