package main

import (
	"fmt"
	"regexp"
)

func main() {
	str := "request_ip[192.168.1.100]"
	pattern := "\\d+.\\d+.\\d+.\\d+"
	if matched, _ := regexp.MatchString(pattern, str); matched {
		fmt.Printf("str=%s contains ip address\n", str)
	}
	// 使用正则的其他功能，例如：替换、查找之类的
	// 要先将正则通过 Compile 方法返回一个 Regexp 对象
	compile, _ := regexp.Compile(pattern)
	allString := compile.FindAllString(str, -1)
	for _, val := range allString {
		fmt.Printf("val = %s\n", val)
	}

}
