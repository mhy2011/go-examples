package main

import (
	"fmt"
	"reflect"
)

func main() {
	student := Student{"20190010001", "张三", 21}

	for i := 0; i < 3; i++ {
		studentType := reflect.TypeOf(student)
		field := studentType.Field(i)
		fmt.Printf("name=%s, type=%v\n", field.Name, field.Type)
	}

}

// 带标签的结构体
type Student struct {
	stuNo string "学号"
	name string "姓名"
	age uint8 "年龄"
}
