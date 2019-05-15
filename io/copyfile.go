package main

import (
	"fmt"
	"io"
	"os"
)

// 文件拷贝示例
func main() {
	srcFile := "file01.txt"
	destFile := "file02.txt"

	_, err := CopyFile(srcFile, destFile)
	if err != nil {
		fmt.Println("copy file error.", err.Error())
	} else {
		fmt.Println("copy file success.")
	}
}

func CopyFile(srcName, destName string) (written int64, err error) {

	srcFile, err := os.Open(srcName)
	if err != nil {
		fmt.Println("open file error.", err.Error())
		return
	}
	defer srcFile.Close()

	destFile, err := os.Create(destName)
	if err != nil {
		fmt.Println("create file error.", err.Error())
		return
	}
	defer destFile.Close()

	i, err := io.Copy(destFile, srcFile)
	return i, err
}
