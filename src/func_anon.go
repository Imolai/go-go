package main

import "fmt"

func main() {
	// Anonymous function
	v := func() {
		fmt.Println("Inside anonymous function")
	}
	v() // call
	func() {
		fmt.Println("Anonymous function gets invoked immediately")
	}()
	func(v int) {
		fmt.Println(v)
	}(42) // Passing Arguments to an Anonymous Function
	g := func(v string) {
		fmt.Println(v)
	}
	func(v string, g func(v string)) {
		g(v)
	}("Passing Anonymous Functions as an Argument", g)
	f1 := func() func(v string) {
		f := func(v string) {
			fmt.Println(v)
		}
		return f
	}
	h := f1()
	h("Returning Anonymous Function from a Function")
}
