package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	i := 40
	res := fibonacci(i)
	cost := time.Now().Sub(start)
	fmt.Printf("fibonacci for %d is %d, cost time is %s\n", i, res, cost)
}

func fibonacci(n int) (res int) {
	if n <= 1 {
		res = 1
	} else {
		res = fibonacci(n - 1) + fibonacci(n - 2)
	}
	return
}
