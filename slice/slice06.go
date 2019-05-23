package main

import "fmt"

func main() {
	s := "你好hello"
	for index, val := range s {
		fmt.Printf("index=%d, val=%c\n", index, val)
	}
	// 获取字符串的某一部分
	fmt.Printf("s[0:5]=%s\n", s[0:5])
}
