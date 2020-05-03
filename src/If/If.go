package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {

	fmt.Println("Simple If ->")

	isGoAwesome := true

	if isGoAwesome {
		fmt.Println("Go is awesome")
	} else {
		fmt.Println("Go should go away!")
	}

	fmt.Println()

	if x := 100; x < Sample() {
		fmt.Println(x, "is lower than Sample")
	} else {
		fmt.Println(x, "is higher than Sample")
	}

	//fmt.Println(x, "Cant access x here")

	fmt.Println()

	fmt.Print("The Go app is currently running on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}

	fmt.Println()

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}

}

func Sample() int {
	return 1337
}
