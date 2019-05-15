package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("请输入内容:")
	input, err := reader.ReadString('\n')
	if err == nil {
		fmt.Printf("输入的内容为:%s\n", input)
	}
}
