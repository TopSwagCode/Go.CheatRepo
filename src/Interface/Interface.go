package main

import (
	"fmt"
	"math"
)

type Geometry interface {
	Area() float64
	Perim() float64
}

type Rect struct {
	width, height float64
}
type Circle struct {
	radius float64
}

type Point struct {
	X int
	Y int
}

func (r Rect) Area() float64 {
	return r.width * r.height
}
func (r Rect) Perim() float64 {
	return 2*r.width + 2*r.height
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c Circle) Perim() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g Geometry) {
	fmt.Println(g)
	fmt.Println(g.Area())
	fmt.Println(g.Perim())
}

func main() {
	r := Rect{width: 3, height: 4}
	c := Circle{radius: 5}

	measure(r)
	measure(c)

	//p := Point{ X:1, Y:2}
	//measure(p) <-- Compiler error. Type does not implement Geometry.
}
