package main

import "fmt"

func main() {
	// 声明时指定初始化值
	nums := [5]int32{1, 3, 5, 7, 9}
	for index, val := range nums {
		fmt.Printf("nums[%d]=%d\n", index, val)
	}
}
