package main

import "fmt"

func main() {
	// 声明切片
	var numbers []int32
	fmt.Printf("len=%d, cap=%d\n", len(numbers), cap(numbers))
	for index, val := range numbers {
		fmt.Printf("index=%d, val=%d\n", index, val)
	}
}
