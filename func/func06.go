package main

import "fmt"

func main() {
	callback(5, 6, Add)
}

func Add(a, b int) {
	fmt.Printf("The sum of %d and %d is: %d\n", a, b, a+b)
}

func callback(x, y int, f func(int, int)) {
	f(x, y)
}
