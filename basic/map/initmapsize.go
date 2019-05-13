package main

import "fmt"

func main() {
	//初始化大小为10的map
	cityMap := make(map[string]string, 10)

	cityMap["bj"] = "北京"
	cityMap["sh"] = "上海"
	cityMap["sz"] = "深圳"

	for key := range cityMap {
		fmt.Println("key =", key, ", value =", cityMap[key])
	}

	fmt.Println("size for map is", len(cityMap))
}
