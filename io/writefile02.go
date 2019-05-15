package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	filename := "test01.txt"
	//以只写追加方式打开文件，如果文件不存在则
	file, e := os.OpenFile(filename, os.O_WRONLY | os.O_APPEND | os.O_CREATE, 0666)
	if e != nil {
		fmt.Println("文件打开失败", e.Error())
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)

	s := "hello world\n"
	for i := 0; i < 10; i++ {
		writer.WriteString(s)
	}
	writer.Flush()

}
