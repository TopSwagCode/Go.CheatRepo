package main

import "fmt"

func main() {

	primes := [6]int{2, 3, 5, 7, 11, 13} // []int{2, 3, 5, 7, 11, 13} are the same
	fmt.Println(primes)

	var slice []int = primes[1:4]
	fmt.Println(slice)

	primes[2] = 1337
	fmt.Println(slice) // The slice still points at the same 3 spots even when array changes

	slice = primes[:2]
	fmt.Println(slice)

	slice = primes[2:]
	fmt.Println(slice)

	fmt.Println(len(primes))
	fmt.Println(cap(primes))

	arr1 := make([]int, 5)
	printSlice("1", arr1)

	arr2 := make([]int, 0, 5)
	printSlice("2", arr2)

	arr3 := arr2[:2]
	printSlice("3", arr3)

	arr4 := arr3[2:5]
	printSlice("4", arr4)

	arr4 = append(arr4, 2)
	printSlice("4", arr4)

	arr4 = append(arr4, 2)
	printSlice("4", arr4)

	arr4 = append(arr4, 2)
	printSlice("4", arr4)

	arr4 = append(arr4, 2)
	printSlice("4", arr4)

	arr4 = append(arr4, 2)
	printSlice("4", arr4)

	for i, v := range arr4 {
		fmt.Printf("primesList**%d = %d\n", i, v)
	}

	for i := range arr4 {
		fmt.Println("primesList**index: %d", i)
	}

	for _, v := range arr4 {
		fmt.Println("primesList**Value: %d", v)
	}

}


func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}