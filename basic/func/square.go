package main

import "fmt"

func main() {
	length := 10
	width := 5

	perimeter, area := getPerimeterAndArea(length, width)
	fmt.Println("perimeter is", perimeter , ",area is", area)

	perimeter2, area2 := getPerimeterAndArea2(length, width)
	fmt.Println("perimeter2 is", perimeter2 , ",area2 is", area2)

}

//求长方形周长和面积
//使用非命名返回值 不建议使用，不明确
func getPerimeterAndArea(length, width int) (int, int) {
	perimeter := 2 * (length + width)
	area := length * width
	return perimeter, area
}

//求长方形周长和面积
//使用命名返回值
func getPerimeterAndArea2(length, width int) (perimeter int, area int) {
	perimeter = 2 * (length + width)
	area = length * width
	return
}
