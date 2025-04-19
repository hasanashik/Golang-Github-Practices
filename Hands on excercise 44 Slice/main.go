package main

import (
	"fmt"
)

func main() {

	mySlice := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	slice1 := mySlice[:5]
	slice2 := mySlice[5:]
	slice3 := mySlice[2:7]
	slice4 := mySlice[1:6]
	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println(slice3)
	fmt.Println(slice4)

	// excercise 45 append slice
	mySlice = append(mySlice, 52)
	fmt.Println(mySlice)
	mySlice = append(mySlice, 53, 54, 55)
	fmt.Println(mySlice)
	y := []int{56, 57, 58, 59, 60}
	mySlice = append(mySlice, y...)
	fmt.Println(mySlice)

}
