package main

import "fmt"

func doMyJob(a int) int {
	a = 5
	return a

}
func main() {
	b := 0
	c := doMyJob(b)

	fmt.Println("b = ", b, " c = ", c)
}
