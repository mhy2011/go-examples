package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("in main()")
	go longWait()
	go shortWait()
	fmt.Println("main ready to sleep")

	// 主线程休眠8S
	time.Sleep(8 * 1e9)
	fmt.Println("finish main()")
}

func longWait() {
	fmt.Println("Begin long Wait......")
	// 休眠5S
	time.Sleep(5 * 1e9)
	fmt.Println("End long Wait......")
}

func shortWait() {
	fmt.Println("Begin short Wait......")
	// 休眠1S
	time.Sleep(1 * 1e9)
	fmt.Println("End short Wait......")
}