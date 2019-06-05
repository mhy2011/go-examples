package main

import "fmt"

func main() {
	i := 20
	fmt.Printf("i = %d , address is %p\n", i, &i)

	// 声明指向int类型的指针
	var ptr *int
	ptr = &i
	fmt.Printf("ptr = %p, *ptr = %d\n", ptr, *ptr)
}
