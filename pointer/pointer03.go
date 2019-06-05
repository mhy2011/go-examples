package main

import "fmt"

func main() {
	s := "Hello"
	fmt.Printf("1.s=%s, addr=%p\n", s, &s)
	s = "World"
	fmt.Printf("2.s=%s, addr=%p\n", s, &s)
	p := &s
	fmt.Printf("3.p=%p\n", p)
	*p = "你好"
	fmt.Printf("4.s=%s, addr=%p\n", s, &s)
}
