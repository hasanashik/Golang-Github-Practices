package main

import (
	"fmt"
)

func main() {
	mySlice := []int{1, 2, 3, 4, 5}

	fmt.Println(foo(mySlice...))
}
func foo(a ...int) int {
	sum := 0
	for _, v := range a {
		sum = sum + v
	}
	return sum
}
