package main

import (
	"fmt"
)

func main() {
	type person struct {
		fname             string
		lname             string
		febIceCreamFlavor []string
	}
	var zaman person
	zaman.fname = "Hasan"
	zaman.lname = "Ashik"
	zaman.febIceCreamFlavor = []string{"vanila", "chocolate"}
	fmt.Printf("%#v", zaman)

}
