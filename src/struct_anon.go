package main

import "fmt"

// --- Anonymous Field in Struct ---

type Polygon struct {
	Sides int
}

func (p *Polygon) NSides() int {
	return p.Sides
}

type Triangle struct {
	Polygon // Anonymous Field
}

// ---------------------------------

func main() {
	pizza := struct {
		name string
	}{
		name: "Pizza",
	}
	fmt.Println(pizza) // prints {Pizza}

	// Using Anonymous Field of Struct to reduce unnecessary function calls.
	t := Triangle{
		Polygon{
			Sides: 3,
		},
	}
	fmt.Println(t.NSides()) // 3
}
