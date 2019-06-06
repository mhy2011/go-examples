package main

import (
	"fmt"
	"sync"
)

func main() {
	info02 := &Info02{name:"test01"}
	fmt.Printf("info02=%v\n", info02)
	Update(info02, "test02")
	fmt.Printf("info02=%v\n", info02)
}

type Info02 struct {
	mx sync.Mutex	//互斥锁
	name string
}

func Update(info *Info02, name string) {
	info.mx.Lock()	//资源加锁
	info.name = name
	info.mx.Unlock()	//释放锁
}