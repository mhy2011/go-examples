package main

import "fmt"

func main() {
	cityMap := make(map[string]string)
	cityMap["bj"] = "北京"
	cityMap["sh"] = "上海"
	cityMap["sz"] = "深圳"

	for key := range cityMap {
		fmt.Println("key =", key, ", value =", cityMap[key])
	}
}
