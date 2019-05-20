package main

import (
	"fmt"
	"time"
)

const LIMIT int = 45

var fibs [LIMIT]uint64

func main() {
	start := time.Now()
	res := fibonacci(LIMIT)
	cost := time.Now().Sub(start)
	fmt.Printf("fibonacci for %d is %d, cost time is %s\n", LIMIT, res, cost)

	start2 := time.Now()
	res2 := fibonacci2(LIMIT)
	cost2 := time.Now().Sub(start2)
	fmt.Printf("fibonacci2 for %d is %d, cost time is %s\n", LIMIT, res2, cost2)

}

func fibonacci(n int) (res uint64) {
	if n <= 2 {
		res = 1
	} else {
		res = fibonacci(n - 1) + fibonacci(n - 2)
	}
	return
}

// 使用内存缓存中间数据
func fibonacci2(n int) (res uint64) {
	index := n -1
	if fibs[index] != 0 {
		res = fibs[index]
		return
	}

	if n <= 2 {
		res = 1
	} else {
		res = fibonacci2(n - 1) + fibonacci2(n - 2)
	}
	//数组存储中间值，避免重复计算
	fibs[index] = res
	return
}
