package main

import (
	"fmt"
)

func main() {

	myMap := make(map[string][]string)
	myMap["bond_james"] = []string{"shaken not, stirred", "martinis", "fast cars"}
	myMap["moneypenny_jenny"] = []string{"intelligence", "literature", "computer science"}
	myMap["no_dr"] = []string{"cats", "ice cream", "sunsets"}

	for k, v := range myMap {
		fmt.Println(k, v)
	}
	fmt.Println("--------------")
	myMap["fleming_ian"] = []string{`steaks`, `cigars`, `espionage`}
	for k, v := range myMap {
		fmt.Println(k, v)
	}
	fmt.Println("--------------")
	delete(myMap, "moneypenny_jenny")
	for k, v := range myMap {
		fmt.Println(k, v)
	}

}
