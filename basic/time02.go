package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()	//获取当前时间
	fmt.Println("now is:", now)
	fmt.Println("utc time is:", now.UTC())//获取对应的UTC时间
	fmt.Printf("year=%v, month=%v, day=%v, hour=%v, minute=%v, second=%v\n",
		now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
	fmt.Println("unix timestamp is:", now.Unix())
	fmt.Println("now format is:", now.Format("2006-01-02 15:04:05"))

	// 字符串转时间
	s := "2019-05-17 10:11:12"
	date, _ := time.Parse("2006-01-02 15:04:05", s)
	fmt.Println("parse date is :", date)
}
