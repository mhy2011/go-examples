package main

import "fmt"

func main() {
	min := min(1, 5, 234, -3, 909, 23)
	fmt.Println("min is", min)
}

func min(a ...int32) int32 {
	if len(a) == 0 {
		return 0
	}
	min := a[0]
	for _, v := range a {
		if v < min {
			min = v
		}
	}
	return min
}
