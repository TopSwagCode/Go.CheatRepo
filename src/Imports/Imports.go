package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	// Time
	fmt.Println("Working with time ->")
	now := time.Now()
	fmt.Println("The time is: ", now)

	fmt.Println()

	fmt.Println("Using time to Sleep thread")
	for i := 0; i < 10; i++ {
		time.Sleep(time.Millisecond * 500)
		fmt.Println(i)
	}

	fmt.Println()

	// Support for different time formats
	fmt.Println("Time formats ->")
	fmt.Println("Kitchen:", now.Format(time.Kitchen))
	fmt.Println("Iso 8601:", now.Format(time.RFC3339))

	fmt.Println()

	// Math
	fmt.Println("Working with Math ->")
	pi := math.Pi
	fmt.Println("Pi:", pi)

	notSoRandom1 := rand.Intn(10)
	notSoRandom2 := rand.Intn(10)
	fmt.Println("Not Random: ", notSoRandom1, notSoRandom2) // I guess this returns 1,7

	rand.Seed(time.Now().UnixNano())
	randomInt1 := rand.Intn(10)
	randomInt2 := rand.Intn(10)
	fmt.Println("Random: ", randomInt1, randomInt2) // I have no idea here

}
