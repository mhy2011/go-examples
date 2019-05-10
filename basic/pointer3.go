package main

import "fmt"

func main() {
	a := 123
	fmt.Println("a address is ", &a)
	doSomething(a)
	doSomething2(&a)
}

func doSomething(a int) {
	b := &a
	fmt.Println("b is ", b)
}

func doSomething2(a *int) {
	c := a
	fmt.Println("c is ", c)
}