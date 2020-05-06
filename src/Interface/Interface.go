package main

import (
	"fmt"
	"math"
)

func main() {
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}

	measure(r)
	measure(c)

	//p := point{ X:1, Y:2}
	//measure(p) <-- Compiler error. Type does not implement Geometry.
}

type geometry interface {
	area() float64
	perimeter() float64
}

type rect struct {
	width, height float64
}
type circle struct {
	radius float64
}

type point struct {
	X int
	Y int
}

func (r rect) area() float64 {
	return r.width * r.height
}
func (r rect) perimeter() float64 {
	return 2*r.width + 2*r.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perimeter())
}

