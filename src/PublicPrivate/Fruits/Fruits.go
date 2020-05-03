package Fruits

import "fmt"

type Fruit interface {
	GetWeight() int
	GetColour() string
}

type pear struct {
	Colour string
	Weight int
}

type Banana struct {
	Colour string
	Weight int
}

func (b Banana) GetWeight() int {
	return b.Weight
}

func (b Banana) GetColour() string {
	return b.Colour
}

type Apple struct {
	Colour string
	weight int
}

func (a Apple) GetWeight() int {
	return a.weight
}

func (a Apple) GetColour() string {
	return a.Colour
}

type Orange struct {
	colour string
	weight int
}

func (o Orange) GetWeight() int {
	return o.weight
}

func (o Orange) GetColour() string {
	return o.colour
}

func PrintFruit(f Fruit) {
	fmt.Printf("Type: %T", f)
	fmt.Print(" Colour: ", f.GetColour())
	fmt.Print(" Weight: ", f.GetWeight())
	fmt.Println()
}
