package main

import "fmt"

func main() {
	rectangle := Rectangle{20, 10}
	fmt.Println("rectangle=", rectangle)

	rect := Rectangle{length: 10, width: 5}
	fmt.Println("rect=", rect)

	fmt.Printf("Area for rect is %d\n", rect.Area())
}

// 定义长方形结构体
type Rectangle struct {
	length uint32 //长度
	width  uint32 //宽度
}

//定义求长方形面积的方法
func (rect Rectangle) Area() (area uint32) {
	area = rect.length * rect.width
	return
}