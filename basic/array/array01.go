package main

import "fmt"

func main() {
	const length = 5
	//声明数组
	var arr [length]int
	for i := 0; i < length; i++ {
		arr[i] = i * 3
	}

	for i := range arr {
		fmt.Println("index =", i, ",value =", arr[i])
	}
}
