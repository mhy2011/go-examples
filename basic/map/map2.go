package main

import "fmt"

func main() {
	cityMap := make(map[string]string)
	cityMap["bj"] = "北京"
	cityMap["sh"] = "上海"
	cityMap["sz"] = "深圳"

	//判断key是否存在
	_, isExist := cityMap["gz"]
	if isExist {
		fmt.Println("key=gz is exist")
	} else {
		fmt.Println("key=gz is not exist")
	}

	// 删除指定的key
	delete(cityMap, "bj")

	for key := range cityMap {
		fmt.Println("key =", key, ", value =", cityMap[key])
	}
}
