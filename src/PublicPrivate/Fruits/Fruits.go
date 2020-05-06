// Package fruits is a collection of fruit struct's used to show private and public fields in Go
package fruits

import "fmt"

// The Fruit interface is to show how interface's work.
// It is used to
type Fruit interface {
	PrintWeight()
	PrintColour()
}

type pear struct {
	Colour string
	Weight int
}

// Banana also know as Banananananana. The fruit of minions!
type Banana struct {
	Colour string
	Weight int
}

// PrintWeight Prints out the weight of the Banana
func (b Banana) PrintWeight() {
	fmt.Println("Weight:",b.Weight)
}

// PrintColour Prints out the Colour of the Banana
func (b Banana) PrintColour() {
	fmt.Println("Colour:",b.Colour)
}

// Apple -> The most expensive fruit. Who knew you needed so many adapters to eat an apple?
type Apple struct {
	Colour string
	weight int
}

// PrintWeight Prints out the weight of the Apple
func (a Apple) PrintWeight() {
	fmt.Println("Weight:",a.weight)
}

// PrintColour Prints out the Colour of the Apple
func (a Apple) PrintColour() {
	fmt.Println("Colour:",a.Colour)
}

// Orange is the best fruit. Because the name describes it's colour.
type Orange struct {
	colour string
	weight int
}

// PrintWeight Prints out the weight of the Orange
func (o Orange) PrintWeight() {
	fmt.Println("Weight:",o.weight)
}

// PrintColour Prints out the Colour of the Orange
func (o Orange) PrintColour() {
	fmt.Println("Colour:",o.colour)
}

// PrintFruit Prints out the Type, Weight and Colour of Fruit
func PrintFruit(f Fruit) {
	fmt.Printf("Type: %T", f)
	fmt.Println()
	f.PrintColour()
	f.PrintWeight()
	fmt.Println()
}
