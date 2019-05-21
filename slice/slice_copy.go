package main

import "fmt"

func main() {
	from := []int{1, 3, 5}
	to := make([]int, 10)
	fmt.Printf("1. to.len=%d, cap=%d, val=%v\n",
		len(to), cap(to), to)

	// 切片复制
	copy(to, from)
	fmt.Printf("2. to.len=%d, cap=%d, val=%v\n",
		len(to), cap(to), to)

	//切片追加
	numbers := []int{1, 2}
	fmt.Printf("1. numbers.address=%p, len=%d, cap=%d, val=%v\n",
		numbers, len(numbers), cap(numbers), numbers)
	numbers = append(numbers, 3, 4, 5)
	fmt.Printf("2. numbers.address=%p, len=%d, cap=%d, val=%v\n",
		numbers, len(numbers), cap(numbers), numbers)
}
