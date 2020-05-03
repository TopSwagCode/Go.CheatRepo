package main

import (
	"TopSwagCode.Go/PublicPrivate/Fruits"
	"encoding/json"
	"fmt"
)

func main() {

	// pear := Fruits.pear {} <-- 404 - Pear not found

	banana := Fruits.Banana{
		Colour: "Yellow",
		Weight: 4,
	}
	Fruits.PrintFruit(banana)

	apple := Fruits.Apple{
		Colour: "Green",
	}
	Fruits.PrintFruit(apple)

	orange := Fruits.Orange{
	}
	Fruits.PrintFruit(orange)

	passionFruit := PassionFruit{
		colour: "Purple",
		weight: 4,
	}
	Fruits.PrintFruit(passionFruit)

	fmt.Println()

	bJson, _ := json.Marshal(banana)
	fmt.Println("Banana:",string(bJson))

	aJson, _ := json.Marshal(apple)
	fmt.Println("Apple",string(aJson))

	oJson, _ := json.Marshal(orange)
	fmt.Println("Orange",string(oJson))

	pfJson, _ := json.Marshal(passionFruit)
	fmt.Println("Passion fruit",string(pfJson))
}


type PassionFruit struct {
	colour string
	weight int
}

func (o PassionFruit) GetWeight() int {
	return o.weight
}

func (o PassionFruit) GetColour() string {
	return o.colour
}