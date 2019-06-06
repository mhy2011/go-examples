package main

import "fmt"

func main() {
	a := 5
	b := 8
	ab := addAB()(a, b)
	fmt.Printf("addAB(%d, %d) = %d\n", a, b, ab)
	ab2 := minusAB()
	fmt.Printf("minusAB(%d, %d) = %d\n", a, b, ab2(a, b))

}

func addAB() func(a, b int) int {
	return func(a, b int) int {
		return a + b
	}
}

func minusAB() func(a, b int) int {
	return func(a, b int) int {
		return a - b
	}
}