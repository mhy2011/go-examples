package main

import "fmt"

func main() {
	person := Person02{"张三", 20}
	fmt.Printf("1.person.address=%p, val=%v\n", &person, person)
	upPerson(person)
	fmt.Printf("2.person.address=%p, val=%v\n", &person, person)
	upPerson2(&person)
	fmt.Printf("3.person.address=%p, val=%v\n", &person, person)
}

func upPerson(p Person02) {
	fmt.Printf("upPerson p.address=%p, val=%v\n", &p, p)
	p.name = "李四"
}

func upPerson2(p *Person02) {
	fmt.Printf("upPerson2 p.address=%p, val=%p\n", &p, p)
	p.name = "赵五"
}


type Person02 struct {
	name string
	age uint8
}