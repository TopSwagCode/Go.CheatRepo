package main

import "fmt"

func main() {
	fmt.Println("Simple Add function ->")
	fmt.Println(addV1(42, 13))
	fmt.Println(addV2(42, 13))

	fmt.Println()

	fmt.Println("Multiple Return's ->")
	a, b := swap("hello", "world")
	fmt.Println(a, b)

	fmt.Println()

	fmt.Println("Named returns ->")
	fmt.Println(addV3(1, 4))

}

func addV1(x int, y int) int {
	return x + y
}

func addV2(x, y int) int { // When two or more consecutive named function parameters share a type, you can omit the type from all but the last.
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func addV3(x, y int) (sum int) {
	sum = x + y
	return // This returns the named return value sum
}
