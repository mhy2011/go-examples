package main

import (
	"fmt"
	"time"
)

func main() {
	// 定义一个通道
	ch := make(chan string)
	go setData(ch)
	go getData(ch)
	time.Sleep(5 * 1e9)
}

func setData(ch chan string) {
	// 往通道发送数据
	ch <- "Hello"
	ch <- "World"
	ch <- "你好"
}

func getData(ch chan string) {
	for ; ; {
		data := <- ch
		fmt.Printf(" data is %s\n", data)
	}
}