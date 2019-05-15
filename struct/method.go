package main

import "fmt"

func main() {
	user := user{"张三", "zhangsan@163.com"}
	sendNotification(user)
}


// 定义一个接口
type notifier interface {
	notify()
}

// 声明一个用户类型
type user struct {
	name string
	email string
}


func (user user) notify() {
	fmt.Printf("sending user email to %s<%s>\n", user.name, user.email)
}

func sendNotification(n notifier) {
	n.notify()
}