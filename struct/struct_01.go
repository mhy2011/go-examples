package main

import "fmt"

func main() {
	//返回t类型的指针
	t := new(T01)
	fmt.Printf("t = %v\n", t)

	var t2 T01
	fmt.Printf("t2 = %v\n", t2)

	t3 := T01{123, "测试"}
	fmt.Printf("t3 = %v\n", t3)
}

// 声明结构体T01
type T01 struct {
	id int32
	name string
}
