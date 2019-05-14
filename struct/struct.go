package main

import "fmt"

func main() {
	person := Person{name: "张三", age: 20}
	fmt.Printf("person.name=%s, person.age=%d\n", person.name, person.age)

	p2 := new(Person)
	p2.name = "李四"
	fmt.Printf("p2.name=%s, p2.age=%d\n", p2.name, p2.age)
}

type Person struct {
	name string
	age  uint8
}
