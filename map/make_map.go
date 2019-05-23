package main

import "fmt"

func main() {
	cityMap := make(map[string]string)
	cityMap["bj"] = "北京"
	cityMap["sh"] = "上海"
	cityMap["sz"] = "深圳"

	// 遍历map
	for key, val := range cityMap {
		fmt.Printf("key=%s, val=%s\n", key, val)
	}

}
