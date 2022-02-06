package main

import "fmt"

// [struct]
var aPoint struct {
	x int
	y int
}

type point struct {
	x int
	y int
}

// User struct is like an object of attributes.
type User struct {
	ID                                  int
	Age                                 int
	Name, FirstName, LastName, Location string
}

// Player struct includes User and additional attribute
type Player struct {
	User
	GameID int
}

func main() {
	// [struct] initialization
	p := point{1, 2}
	q := &point{1, 2} // has type *point
	r := point{x: 1}  // y:0 is implicit
	s := point{}      // x:0 and y:0
	fmt.Printf("point{1, 2}: %v,\n", p)
	fmt.Printf("&point{1, 2}: %v,\n", q)
	fmt.Printf("point{x: 1}: %v,\n", r)
	fmt.Printf("point{}: %v\n", s)
	me := Player{}
	me.ID = 47
	me.Name = "Gábor"
	me.Location = "Hungary"
	fmt.Println(me)
}
