package main

import "fmt"

func main() {
	rectangle := Rectangle{20, 10}
	fmt.Println("rectangle=", rectangle)

	rect := Rectangle{length: 10, width: 5}
	fmt.Println("rect=", rect)
}

// 定义长方形结构体
type Rectangle struct {
	length uint32 //长度
	width  uint32 //宽度
}
