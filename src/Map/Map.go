package main

import "fmt"

func main() {
	var m map[string]Point
	m = make(map[string]Point)

	m["test"] = Point{1, 1}

	fmt.Println(m["test"])

	m2 := map[string]Point{
		"Test": {40, -74},
		"Ing":  {37, -122},
	}

	fmt.Println(m2)

	// Check if element exists
	elem, ok := m2["test"]

	if ok {
		fmt.Println(elem)
	}

	// Make a map containing
	m3 := map[string]interface{}{"a": "apple", "b": 2, "c": []interface{}{"foo", 2, "bar", false, map[string]interface{}{"baz": "bat", "moreFoo": 7}}}

	fmt.Println(m3)
}

type Point struct {
	X int
	Y int
}
