package main

import "fmt"

func main() {
	cityMap := make(map[string]string)
	cityMap["bj"] = "北京"
	cityMap["sh"] = "上海"
	cityMap["sz"] = "深圳"

	// 检查key是否存在
	key := "jn"
	val, ok := cityMap[key]
	if !ok {
		fmt.Printf("key=%s is not exist. val=%v\n", key, val)
	}

	key2 := "sh"
	val, ok = cityMap[key2]
	if ok {
		fmt.Printf("key=%s is exist, val=%s\n", key2, val)
	}

	// 删除指定的key
	delete(cityMap, key2)
	// 遍历map
	for key, val := range cityMap {
		fmt.Printf("cityMap.key=%s, val=%s\n", key, val)
	}
}
