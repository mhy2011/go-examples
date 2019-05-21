package main

import "fmt"

func main() {
	numbers := [5]int{1, 2, 3, 4, 5}
	//sliceSum(numbers)	//不能直接传数组，要传切片形式
	res := sliceSum(numbers[:])
	fmt.Println("res is:", res)
}

// 将切片传递给函数
// 求切片数据和
func sliceSum(nums []int) (res int) {
	for _, val := range nums {
		res += val
	}
	return
}
