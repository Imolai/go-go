package main

import "fmt"

func main() {
	// [if]
	answer := 42
	if answer == 42 { // [==]
		fmt.Println("The Answer to the Ultimate Question of Life, the Universe, and Everything...")
	}
	foo := func() interface{} {
		return nil
	}
	// Like for, the if statement can start with a short statement.
	if err := foo(); err != nil { // [!=]
		panic(err)
	}
	// Nested ifs can be avoided by chaining conditions logically.
	question := "What do you get if you multiply six by nine?"
	if question == "What do you get if you multiply six by nine?" && answer == 42 {
		fmt.Println("Don't panic!")
	}
}
