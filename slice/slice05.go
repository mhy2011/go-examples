package main

import "fmt"

func main() {
	numbers := make([]int, 0, 10)
	fmt.Printf("numbers.len=%d, cap=%d, val=%v\n", len(numbers), cap(numbers), numbers)
	//下标越界，因为切片长度为0
	//numbers[3] = 1
	for i := 0; i < cap(numbers); i++ {
		numbers = numbers[0:i+1]	//切片重组
		fmt.Printf("i=%d, numbers.len=%d, cap=%d, val=%v\n",
			i, len(numbers), cap(numbers), numbers)
	}
}
