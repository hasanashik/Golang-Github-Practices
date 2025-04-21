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

	// second person
	farid := person{fname: "farid", lname: "hossain", febIceCreamFlavor: []string{"mango", "chocolate"}}
	fmt.Printf("%v", farid)
	fmt.Println("\n")
	myMap := map[string]person{
		zaman.lname: zaman,
		farid.lname: farid,
	}
	for _, v := range myMap {
		fmt.Println(v)
	}

}
