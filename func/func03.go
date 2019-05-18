package main

import "fmt"

func main() {
	length := 20
	width := 10

	//丢弃面积，只要周长
	perimeter, _ := getPerimeterAndArea(length, width)
	fmt.Println("perimeter is:", perimeter)
}

//求长方形周长和面积
//使用命名返回值
func getPerimeterAndArea(length, width int) (perimeter int, area int) {
	perimeter = 2 * (length + width)
	area = length * width
	return
}
