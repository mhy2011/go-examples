package main

import "fmt"

func main() {
	ch := make(chan int)
	go pump(ch)
	fmt.Println(<-ch)
}

func pump(ch chan int) {
	for i := 0; ; i++ {
		//向通道中写数据
		ch <- i
	}
}
