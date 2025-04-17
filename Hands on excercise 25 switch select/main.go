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
		fmt.Print("between 0 and 100\n")
	case x <= 200 && x > 100:
		fmt.Print("between 101 and 200\n")
	case x <= 250 && x > 200:
		fmt.Print("between 201 and 250\n")
	default:
		fmt.Print("x is out of expected range!\n")
	}

	// Create channels for each range
	range1 := make(chan string, 1)
	range2 := make(chan string, 1)
	range3 := make(chan string, 1)
	// Send value to the appropriate channel
	switch {
	case x <= 100:
		range1 <- "between 0 and 100"
	case x <= 200:
		range2 <- "between 101 and 200"
	case x <= 250:
		range3 <- "between 201 and 250"
	}

	// Use select to receive from one of the channels
	select {
	case msg := <-range1:
		fmt.Println(msg)
	case msg := <-range2:
		fmt.Println(msg)
	case msg := <-range3:
		fmt.Println(msg)
	}
}
