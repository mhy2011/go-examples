package main

import "fmt"

func main() {
	callback(5, 10, Add)
}

func Add(a, b int) {
	fmt.Printf("%d + %d = %d\n", a, b, a+b)
}

func callback(x int, y int, f func(int, int)) {
	fmt.Printf("callback is execute ......\n", )
	f(x, y)
}
