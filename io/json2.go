package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var jsonStr = "{\"Name\":\"张三\",\"StuNo\":\"20190010001\",\"Age\":21}"

	var stu Student2

	e := json.Unmarshal([]byte(jsonStr), &stu)
	if e != nil {
		fmt.Println("error .", e.Error())
		return
	}
	fmt.Println("stu is : ", stu)
}

type Student2 struct {
	Name string
	StuNo string
	Age uint8
}
