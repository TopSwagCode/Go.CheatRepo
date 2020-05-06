package main

import "fmt"

func main() {
	numberValue := 42
	numberPointer := &numberValue

	fmt.Println("number", numberValue)
	fmt.Println("pointer", numberPointer)

	fmt.Println()
	fmt.Println("Set number")
	numberValue = 32

	fmt.Println("number", numberValue)
	fmt.Println("pointer", numberPointer)

	fmt.Println()
	fmt.Println("Set pointer value")
	*numberPointer = 52

	fmt.Println(numberValue)
	fmt.Println("pointer", numberPointer)

	fmt.Println()

	fmt.Println(point{1, 2})

	p := point{1, 2}
	fmt.Println(p.X)

	pointPointer := &p
	pointPointer.X = 3
	fmt.Println(p.X)

	fmt.Println()

	// Go is pass by value. Example shown below
	answer := 42

	fmt.Println("answer befor", &answer)
	fmt.Println("answer befor", answer)
	foo(answer)
	fmt.Println("answer after", &answer)
	fmt.Println("answer after", answer)

	fmt.Println()

	// Now with pointers
	fmt.Println("answer befor", &answer)
	fmt.Println("answer befor", answer)
	fooPointer(&answer)
	fmt.Println("answer after", &answer)
	fmt.Println("answer after", answer)

	/*
		Can pointers negatively affect performance?
		Absolutely. There are two major considerations here:
		Dereferencing pointers isn’t free. It’s not a huge cost, but it can add up.
		Sharing data via pointers will likely cause the data to be placed in the “heap.”
		The heap is a section of memory for data that lives longer than a single function call.
		There is overhead to adding data to the heap and heap data can only be cleaned up by the garbage collector.
		The more data in the heap, the more work the garbage collector has to do, and the more impact it’ll have on your application.
	*/
}

type point struct {
	X int
	Y int
}

func foo(x int) {
	x = x * x
	fmt.Println("answer updat", &x)
	fmt.Println("answer updat", x)
}

func fooPointer(x *int) {
	*x = *x * *x
	fmt.Println("answer updat", x)
	fmt.Println("answer updat", *x)
}
