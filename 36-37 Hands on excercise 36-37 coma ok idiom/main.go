package main

import (
	"fmt"
)

func main() {
	m := map[string]int{
		"James":      42,
		"Moneypenny": 32,
		"Q":          0,
	}
	for k, v := range m {
		fmt.Println(k, v)
	}
	age := m["James"]
	fmt.Println(age)
	value, ok := m["Q"]
	if ok {
		fmt.Println("Q is fount and ther value is ", value)
	} else {
		fmt.Println("Q key is not found in m.")
	}

}
