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

	// excercise 46 delete from slice
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	x = append(x[:3], x[6:]...)
	fmt.Println(x)

	// using make
	states := []string{` Alabama`, ` Alaska`, ` Arizona`, ` Arkansas`, ` California`, ` Colorado`, ` Connecticut`, `Delaware`, ` Florida`, ` Georgia`, ` Hawaii`, ` Idaho`, ` Illinois`, ` Indiana`, ` Iowa`, ` Kansas`, `Kentucky`, ` Louisiana`, ` Maine`, ` Maryland`, ` Massachusetts`, ` Michigan`, ` Minnesota`, `Mississippi`, ` Missouri`, ` Montana`, ` Nebraska`, ` Nevada`, ` New Hampshire`, ` New Jersey`, ` New Mexico`, ` New York`, ` North Carolina`, ` North Dakota`, ` Ohio`, ` Oklahoma`, ` Oregon`, ` Pennsylvania`, ` Rhode Island`, ` South Carolina`, ` South Dakota`, ` Tennessee`, ` Texas`, `Utah`, ` Vermont`, ` Virginia`, ` Washington`, ` West Virginia`, ` Wisconsin`, ` Wyoming`}
	statesMake := make([]string, 0, 50)
	fmt.Println("Length: ", len(statesMake))
	fmt.Println("Capacity: ", cap(statesMake))
	statesMake = append(statesMake, states[:]...)
	fmt.Println("Length: ", len(statesMake))
	fmt.Println("Capacity: ", cap(statesMake))
	fmt.Println(statesMake)

}
