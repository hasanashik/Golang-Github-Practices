package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	x := rand.IntN(250)
	fmt.Print("x = ", x, "\n")
	if x <= 100 {
		fmt.Print("between 0 and 100")
	} else if x <= 200 && x > 100 {
		fmt.Print("between 101 and 200")
	} else {
		fmt.Print("between 201 and 250")
	}
}
