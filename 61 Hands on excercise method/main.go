package main

import (
	"fmt"
)

type person struct {
	name string
	age  int
}

func (p person) speak() {
	fmt.Println("I am", p.name, "and my age is", p.age)
}
func main() {

	zaman := person{
		name: "zaman",
		age:  30,
	}
	zaman.speak()
}
