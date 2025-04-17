package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	x := rand.IntN(250)
	fmt.Print("x = ", x, "\n")
	switch {
	case x <= 100:
		fmt.Print("between 0 and 100")
	case x <= 200 && x > 100:
		fmt.Print("between 101 and 200")
	case x <= 250 && x > 200:
		fmt.Print("between 201 and 250")
	default:
		fmt.Print("x is out of expected range!")
	}
}
