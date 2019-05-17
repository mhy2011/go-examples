package main

import "fmt"

var PI float32

func main() {
	fmt.Println("exec main func. PI =", PI)
}

func init() {
	fmt.Println("exec init func ......")
	PI = 3.14
}