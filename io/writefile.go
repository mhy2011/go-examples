package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	filename := "file01.txt"
	s := "你好abc"
	bytes := [] byte(s)

	err := ioutil.WriteFile(filename, bytes, 0777)
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("文件写成功")
	}
}
