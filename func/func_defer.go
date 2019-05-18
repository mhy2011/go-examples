package main

import "fmt"

func main() {
	fmt.Println("1. This is first to print")
	defer fmt.Println("2. This is last to print")
	fmt.Println("3. This is second to print")
}
