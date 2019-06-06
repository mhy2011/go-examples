package main

import "fmt"

func main() {
	f := NewFile02(1, "test.txt")
	fmt.Printf("f = %v\n", f)
}

type File02 struct {
	fd   int    //文件描述符
	name string //文件名称
}

//自定义构造方法
func NewFile02(fd int, name string) *File02 {
	return &File02{fd, name}
}