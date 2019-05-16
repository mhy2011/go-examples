package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Name string
	StuNo string	//大写开头才能在json里显示
	Age uint8
}

func main() {
	student := Student{"张三", "20190010001", 21}
	stuJson, _ := json.Marshal(student)
	fmt.Printf("\nstuJson is : %s\n", string(stuJson))
}