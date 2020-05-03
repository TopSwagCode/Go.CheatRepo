package main

import "fmt"

func main() {

	p1 := Point{1, 2} // has type Point
	p2 := Point{X: 1} // Y:0 is implicit
	p3 := Point{}     // X:0 and Y:0

	fmt.Println(p1, p2, p3)

	// Go doesn't support constructors, but constructor-like factory functions are easy to implement:
	p4 := New(1, 2)
	fmt.Println(p4)
}

type Point struct {
	X int
	Y int
}

func New(x, y int) Point {
	var p Point

	if x != 0 {
		p.X = x
	} else {
		p.X = 15
	}

	if y != 0 {
		p.Y = y
	} else {
		p.Y = 15
	}

	return p
}
