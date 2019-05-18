package main

import "fmt"

func main() {
	i := 10
	res := fibonacci(i)
	fmt.Printf("fibonacci for %d is %d\n", i, res)
}

func fibonacci(n int) (res int) {
	if n <= 1 {
		res = 1
	} else {
		res = fibonacci(n - 1) + fibonacci(n - 2)
	}
	return
}
