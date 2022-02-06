package main

import "fmt"

func add(x int, y int) int {
	// [return]
	return x + y // [+]
}

// Can return more values.
func sort(x, y int) (int, int) {
	if x < y {
		return x, y // [,]
	}
	return y, x
}

/*
	Multiple return values can be named and act just like variables. If the
	result parameters are named, a return statement without arguments returns the
	current values of the results.
*/
func sort2(x, y int) (first, second int) {
	if x < y {
		first, second = x, y // [,]
	} else {
		first, second = y, x
	}
	return
}

func main() {
	fmt.Println("add(42, 13): ", add(42, 13))
	fmt.Println("mymath(12.0, 33.0): ")
	a, b := sort(42, 13)
	fmt.Println("sort(42, 13): ", a, b)
	a, b = sort2(42, 13)
	fmt.Println("sort2(42, 13): ", a, b)
}
