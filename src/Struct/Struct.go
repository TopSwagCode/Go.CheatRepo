package main

import "fmt"

func main() {

	p1 := point{1, 2} // has type Point
	p2 := point{X: 1} // Y:0 is implicit
	p3 := point{}     // X:0 and Y:0

	fmt.Println(p1, p2, p3)

	// Go doesn't support constructors, but constructor-like factory functions are easy to implement:
	p4 := new(1, 2)
	fmt.Println(p4)
}

type point struct {
	X int
	Y int
}

func new(x, y int) point {
	var p point

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
