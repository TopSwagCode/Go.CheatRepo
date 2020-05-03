package main

import "fmt"

func main() {

	sum1 := 0
	for i := 0; i < 10; i++ {
		sum1 += i
	}
	fmt.Println(sum1)

	fmt.Println()

	sum2 := 1
	for sum2 < 1000 { // The init and post statements are optional.
		sum2 += sum2
	}
	fmt.Println(sum2)

	fmt.Println()

	sum3 := 1
	for sum3 < 1000 {
		sum3 += sum3
	}
	fmt.Println(sum3)

	/*
		running := true
		for running{
			fmt.Println("running")
		}
	*/
}
