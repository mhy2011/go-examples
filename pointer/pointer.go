package main

import "fmt"

func main() {
	i := 20
	fmt.Printf("value for i is:%d, mem address is:%p\n", i, &i)
	// 声明一个int类型指针变量
	var p *int
	// p的值实际上存储的是变量i的内存地址
	p = &i
	fmt.Printf("value for p is:%p, mem address is:%v, point value is:%d\n",
		p, &p, *p)
	// 变更指针所指向变量的值为10
	*p = 10
	fmt.Printf("value for p is:%p, p point value is:%d, i=%d\n", p, *p, i)
}
