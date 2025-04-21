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
	fmt.Printf("%v", zaman)
	fmt.Println("\n")

	myMap := make(map[string][]string)
	myMap[zaman.lname] = []string{"Hasan", "Ashik", "vanila", "chocolate"}
	for _, v := range myMap {
		fmt.Println(v)
	}

}
