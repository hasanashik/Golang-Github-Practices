package main

import (
	"fmt"
	"math"
)

type square struct {
	length float64
	width  float64
}
type circle struct {
	radius float64
}

func (s square) area() float64 {
	return s.length * s.width
}
func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

type shape interface {
	area() float64
}

func info(s shape) float64 {
	return s.area()
}
func main() {
	shape1 := square{
		length: 10.0,
		width:  12.0,
	}
	fmt.Println(info(shape1))

	shape2 := circle{
		radius: 5.0,
	}
	fmt.Println(shape2.area())

}
