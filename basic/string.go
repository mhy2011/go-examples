package main

import (
	"fmt"
	"strings"
)

func main() {
	strEn := "abcd"
	strZh := "你好"

	fmt.Printf("strEn's length is %d, strZh's length is %d\n", len(strEn), len(strZh))
	upperStr := strings.ToUpper(strEn)
	fmt.Printf("upperStr is %s\n", upperStr)
}
