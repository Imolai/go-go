// Package main provides possibility to see Go documentation.
package main

import "fmt"

// Version represents the version number of the package.
const Version = 1.0

// Help returns nothing, but prints the usage as side effect.
func Help() {
	fmt.Println("My main package. Created by me.")
	fmt.Printf("Version: %.1f\n", Version)
}

func main() {
	/*
		General comments start with the character sequence slash-star and stop with
		the first subsequent character sequence star-slash.
		To document a type, variable, constant, function, or even a package, write
		a regular comment directly preceding its declaration, with no intervening
		blank line.
	*/
	// Line comments start with the character sequence slash-slash and stop at
	// the end of the line.
	Help()
}
