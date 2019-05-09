package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()	//获取当前时间
	fmt.Println("now is ", now)
	fmt.Printf("curYear=%d, curMonth=%d, curDay=%d\n", now.Year(), now.Month(), now.Day())
	now.Day()
}
