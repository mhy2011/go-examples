package main

import "fmt"

const LENGTH int = 5

func main() {
	//声明int类型的数组，长度必须为常量or魔数
	var ary [LENGTH]int
	//默认情况下,int类型数组初始化值为0
	for index, val := range ary {
		fmt.Printf("1. ary[%d] = %d\n", index, val)
	}
	//指定数组中数据值
	for i := 0; i < LENGTH; i++ {
		ary[i] = i * 5
	}

	for index, val := range ary {
		fmt.Printf("2. ary[%d] = %d\n", index, val)
	}
}
