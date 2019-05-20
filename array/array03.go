package main

import "fmt"

func main() {
	// 声明二维数组5行3列
	var ary [5][3]int
	// 设置数组值
	ary[1][2] = 10
	ary[2][0] = 3

	for index, val := range ary {
		fmt.Printf("index=%d, val=%v\n", index, val)
	}

	//遍历二维数组
	for i := 0; i < 5; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("ary[%d][%d] = %v\n", i, j, ary[i][j])
		}
	}

}
