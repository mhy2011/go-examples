package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	filename := "hello.txt"
	file, e := os.Open(filename) //只读方式打开文件
	if e != nil {
		fmt.Printf("文件不存在or没有权限\n", )
		return
	}
	defer file.Close()	//关闭文件

	reader := bufio.NewReader(file)
	for {
		s, e := reader.ReadString('\n')
		fmt.Printf(s) //输出文件一行的内容
		if e == io.EOF {
			return
		}
	}
}
