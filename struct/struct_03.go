package main

import (
	"fmt"
	"go-examples/util"
)

func main() {
	x, y := 10, 20
	sum := util.SumXY(x, y)
	fmt.Printf("SumXY(%d, %d) = %d\n", x, y, sum)
}
