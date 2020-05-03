package main

import "fmt"

func main() {

	p1 := Point{1, 2} // has type Point
	p2 := Point{X: 1} // Y:0 is implicit
	p3 := Point{}     // X:0 and Y:0

	fmt.Println(p1, p2, p3)

	// Go doesn't support constructors, but constructor-like factory functions are easy to implement:
	p4 := NewPoint(1,2)
	fmt.Println(p4)
}


type Point struct {
	X int
	Y int
}

func NewPoint(x, y int) *Point {
	p := new(Point)
	p.X = x
	p.Y = y

	return p
}
