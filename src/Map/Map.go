package main

import "fmt"

func main() {
	var m map[string]point
	m = make(map[string]point)

	m["test"] = point{1, 1}

	fmt.Println(m["test"])

	m2 := map[string]point{
		"Test": {40, -74},
		"Ing":  {37, -122},
	}

	fmt.Println(m2)

	// Check if element exists
	elem, ok := m2["test"]

	if ok {
		fmt.Println(elem)
	}

	// Make a map containing different objects.
	// Using the empty interface is generally discouraged
	// Don't do this at home :)
	m3 := map[string]interface{}{"a": "apple", "b": 2, "c": []interface{}{"foo", 2, "bar", false, map[string]interface{}{"baz": "bat", "moreFoo": 7}}}

	fmt.Println(m3)
}

type point struct {
	X int
	Y int
}
