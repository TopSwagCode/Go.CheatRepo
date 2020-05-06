package main

import "fmt"

func main() {

	primes := [6]int{2, 3, 5, 7, 11, 13} // []int{2, 3, 5, 7, 11, 13} are the same
	fmt.Println("Primes:", primes)

	var number int = primes[0]
	fmt.Println("Index 0:", number)

	var slice []int = primes[1:4]
	fmt.Println("Slice 1:4 :", slice)

	primes[2] = 1337
	fmt.Println("Primes:", primes)
	fmt.Println("Slice:", slice) // The slice still points at the same 3 spots even when array changes

	slice = primes[:2]
	fmt.Println("Slice :2 :", slice)

	slice = primes[2:]
	fmt.Println("Slice 2: :", slice)

	// To get the length and cap of a array
	fmt.Println("Length of primes:", len(primes))
	fmt.Println("Cap of primes:", cap(primes))

	fmt.Println()

	fmt.Println("Making an Slice with room for 5 elements and adding elements to beyond Cap to see what happens")

	slice1 := make([]int, 5)
	printSlice(slice1)

	slice2 := make([]int, 0, 5)
	printSlice(slice2)

	for i := 0; i < 10; i++ {
		slice2 = append(slice2, i*i)
		printSlice(slice2)
	}

	fmt.Println()
	fmt.Println()

	fmt.Println("Showing of 3 Ways of working with all elements in an Array.")
	fmt.Println()

	fmt.Println("Index and value:")
	for i, v := range slice2 {
		fmt.Printf("slice**%d = %d\n", i, v)
	}

	fmt.Println()
	fmt.Println("Index only:")
	for i := range slice2 {
		fmt.Println("slice**index: ", i)
	}

	fmt.Println()
	fmt.Println("Value only:")
	for _, v := range slice2 {
		fmt.Println("slice**Value: ", v)
	}

	fmt.Println()
	fmt.Println("Subset of Value only:")
	for _, v := range slice2[1:5] {
		fmt.Println("slice**Value: ", v)
	}

}

func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d %v\n",
		len(x), cap(x), x)
}
