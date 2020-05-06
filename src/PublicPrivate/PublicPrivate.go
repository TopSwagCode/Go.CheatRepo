package main

import (
	"./fruits"
	"encoding/json"
	"fmt"
)

func main() {

	fmt.Println("Using fruits interface to print the different types of fruit and see what fields get left out.")
	// pear := Fruits.pear {} <-- 404 - Pear not found
	// Because pear is in lowercase and therefor not exported

	banana := fruits.Banana{
		Colour: "Yellow",
		Weight: 4,
	}
	fruits.PrintFruit(banana)

	apple := fruits.Apple{
		Colour: "Green",
	}
	fruits.PrintFruit(apple)

	orange := fruits.Orange{}
	fruits.PrintFruit(orange)

	passionFruit := passionFruit{
		colour: "Purple",
		weight: 4,
	}
	fruits.PrintFruit(passionFruit)

	fmt.Println()
	fmt.Println("Using built in Json package to get json strings from objects, so we can see fields getting left out.")

	bJson, e := json.Marshal(banana)
	if e != nil {
		fmt.Println(e)
	}
	fmt.Println("Banana:", string(bJson))

	aJson, e := json.Marshal(apple)
	if e != nil {
		fmt.Println(e)
	}
	fmt.Println("Apple", string(aJson))

	oJson, e := json.Marshal(orange)
	if e != nil {
		fmt.Println(e)
	}
	fmt.Println("Orange", string(oJson))

	pfJson, e := json.Marshal(passionFruit)
	if e != nil {
		fmt.Println(e)
	}
	fmt.Println("Passion fruit", string(pfJson))
}

type passionFruit struct {
	colour string
	weight int
}

// Prints out the weight of the Passion fruit
func (pf passionFruit) PrintWeight() {
	fmt.Println("Weight:", pf.weight)
}

// Prints out the Colour of the Passion fruit
func (pf passionFruit) PrintColour() {
	fmt.Println("Colour:", pf.colour)
}
