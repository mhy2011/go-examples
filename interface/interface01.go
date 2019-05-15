package main

import "fmt"

const PI float32 = 3.14

func main() {
	square := Square{3}
	circle := Circle{4}

	shapers := [] Shaper{square, circle}
	for _, shaper := range shapers {
		fmt.Println("area is :", shaper.Area())
	}

}

// 形状接口
type Shaper interface {
	Area() float32
}

type Square struct {
	side float32 "边长"
}

type Circle struct {
	radius float32 "半径"
}

// 正方形面积
func (square Square) Area() float32 {
	return square.side * square.side
}

func (circle Circle) Area() float32 {
	return PI * circle.radius * circle.radius
}