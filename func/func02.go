package main

import "fmt"

func main() {
	i := 50
	fmt.Println("before value for i is :", i)
	do(i)
	fmt.Println("after call do value for i is :", i)
	do2(&i)
	fmt.Println("after call do2 value for i is :", i)

}

func do(i int) {
	i = 100
}

func do2(i *int) {
	*i = 200
}
