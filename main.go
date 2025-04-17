package main

import (
	"fmt"

	"github.com/hasanashik/puppy"
)

var color string = "Red"

const brand string = "Tesla"

func main() {
	model := "Model 3"
	fmt.Println("hello world")
	fmt.Println("Puppy barks: ", puppy.Bark())
	fmt.Println(puppy.BigBark())
	fmt.Printf("This is a %v car with color %v and the model is %v", brand, color, model)
}
