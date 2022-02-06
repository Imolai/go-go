package main

import (
	"fmt"
	"time"
)

// user struct is like an OBJECT of attributes.
type user struct {
	id                                  int
	age                                 int
	name, firstName, lastName, location string
}

// [interface]
// Namer [interface] contains one method definition.
// An interface type is defined by a set of methods.
type Namer interface {
	Name() string
}

// Name <METHOD> prints the user's name.
// Interfaces are satisfied implicitly.
func (u *user) Name() string {
	return fmt.Sprintf("%s %s", u.firstName, u.lastName)
}

/*
	Stringer [interface] is defined by the fmt package.
	This is one of the most ubiquitous interfaces.
	A STRINGER is a type that can describe itself as a string.
*/
// Stringer interface defines String() method.
type Stringer interface {
	String() string
}

func (u *user) String() string {
	return fmt.Sprintf("--- USER ---\nname:\t\t%v\nid:\t\t%v\nage:\t\t%v\nlocation:\t%v\n------------", u.Name(), u.id, u.age, u.location)
}

// The ERROR type is a built-in interface similar to fmt.Stringer:
type error interface {
	Error() string
}
type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}
func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

func main() {
	me := new(user)
	me.id = 1
	me.age = 127
	me.firstName = "Gábor"
	me.lastName = "Imolai"
	me.location = "Hungary"
	fmt.Println(me) // Print via Stringer.
}
