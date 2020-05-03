package main

import "fmt"

var c, python, java bool

func main() {

	// One way
	var i int = 1
	var j float64 = 2
	fmt.Println("i:",i, "j:", j)
	j = j / 3.14
	fmt.Println("i:",i, "j:", j)

	// Or another.
	x := 1
	y := 2.0
	fmt.Println("x:",x, "y:",y)
	y = y / 3.14
	fmt.Println("x:",x, "y:",y)

	/*
	 I'm gonna find ya
	 I'm gonna get ya, get ya, get ya, get ya
	 One way or another, I'm gonna win ya
	 I'm gonna get ya, get ya ,get ya, get ya
	 One way or another, I'm gonna see ya
	 I'm gonna meet ya, meet ya, meet ya, meet ya
	 One day, maybe next week, I'm gonna meet ya
	 I'm gonna meet ya, I'll meet ya
	 */

	// Defaults
	var q bool
	var w int
	var e string

	fmt.Println("Defaults:",q,w,e)

	fmt.Println(c,python,java)

	fmt.Println()

	// Alternative initializers
	var isFriday, isMonday, javaIsAwesome = true, false, "no!"
	fmt.Println(x, y, isFriday, isMonday, javaIsAwesome)

	var a, s int = 1, 2
	fmt.Println(a,s)

	f, g := 3, 4
	fmt.Println(f, g)
}

