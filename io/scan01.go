package main

import "fmt"

func main() {
	var name string
	fmt.Println("请输入全名:")
	fmt.Scanln(&name)
	fmt.Println("姓名为:", name)

	// 接收多个参数
	fmt.Println("请输入性别和年龄:")
	var gender string
	var age uint8
	fmt.Scanln(&gender, &age)
	fmt.Println("性别:", gender, ", 年龄:", age)
}
