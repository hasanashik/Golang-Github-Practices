package main

import (
	"fmt"

	"github.com/hasanashik/puppy"
)

func main() {
	fmt.Println("hello world")
	fmt.Println("Puppy barks: ", puppy.Bark())
	fmt.Println("Grown up puppy barks: ", puppy.Barks())
}
