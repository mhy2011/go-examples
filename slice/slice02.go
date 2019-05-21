package main

import "fmt"

func main() {

	//定义并初始化一个数组
	numbers := [5]int{1, 3, 5, 7, 9}
	fmt.Printf("numbers.address=%p, len=%d, val=%v\n", &numbers, len(numbers), numbers)
	slice := numbers[1:3]
	// cap(slice)=slice[0]到numbers末尾的长度
	fmt.Printf("slice.address=%p, len=%d, cap=%d, val=%v\n",
		slice, len(slice), cap(slice), slice)
	// 从index=2到数组结尾
	slice2 := numbers[2:]
	fmt.Printf("slice2.address=%p, len=%d, cap=%d, val=%v\n",
		slice2, len(slice2), cap(slice2), slice2)
	//切片是对数组的引用，修改切片内容影响数组内容
	slice2[0] = 123
	fmt.Printf("numbers.address=%p, len=%d, val=%v\n", &numbers, len(numbers), numbers)
	// 长度溢出
	slice2[10] = 100
	fmt.Printf("slice2.address=%p, len=%d, cap=%d, val=%v\n",
		slice2, len(slice2), cap(slice2), slice2)
}
