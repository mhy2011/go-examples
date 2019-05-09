package main

import "fmt"

func main() {
	data := 20
	fmt.Printf("data value is %d, address is %x\n", data, &data)
	//指向整形的指针
	var p *int
	p = &data
	fmt.Println("p is ", p , ",value is", *p)
	//指向字符串的指针
	str := "abcd"
	var p2 *string
	p2 = &str
	fmt.Println("p2 is", p2, ",value is", *p2)
	// 直接定义指针
	p3 := &str
	fmt.Println("p3 is", p3, ",value is", *p3)
}
