package main

import (
	"fmt"
	"time"
)

func main() {

	// Timer example using Defer
	stop := StartTimer("Sample Timer")
	defer stop()

	defer fmt.Println("Have a nice day! :) ")

	DeferInsideFunction()

	// What will be printed here? Hello or World?
	m := "Hello"
	defer fmt.Println(m)
	m = "world"

	// What happens here ?
	for i := 0; i < 10; i++ {
		defer func(j int) {
			fmt.Println(j)
		}(i)
	}

	// ?
	for i := 0; i < 10; i++ {
		defer func() {
			fmt.Println(i)
		}()
	}

	// ?
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	time.Sleep(1 * time.Second)

	fmt.Println("done")

}

func StartTimer(name string) func() {
	t := time.Now()
	fmt.Println(name, "started")
	return func() {
		d := time.Now().Sub(t)
		fmt.Println(name, "took", d)
	}
}

func DeferInsideFunction() {
	fmt.Println("Working in function")
	defer fmt.Println("Defer inside function")
	fmt.Println("Done working inside function")
}
