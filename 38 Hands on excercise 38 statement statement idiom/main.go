package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {

	for counter := 0; counter < 100; counter++ {

		if x := rand.IntN(5); x == 3 {
			fmt.Printf("counter = %v x is 3\n", counter)
		}

	}

}
