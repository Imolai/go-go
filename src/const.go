package main

import "fmt"

// [const] Constants

// Pi exported constant, one value per line.
const Pi = 3.14 // const cannot be declared using the := syntax

// More constants by one const.
const (
	truth    = false           // boolean constant
	lf       = '\n'            // rune constant
	statusOK = 200             // integer constant
	pi       = 3.1415965359    // floating-point constant
	comp     = 1.e+0i          // complex constant
	greeting = "Hello, world!" // string constant
)

const (
	a = 0
	b = 1
	c = 2
)

// The same with auto increment IOTA.
const (
	d = iota // 0
	e        // 1
	f        // 2
)

// Iota can also start from non-zero number - iota expressions can also be used
// to start iota from any number
const (
	g = iota + 10 // 10
	h             // 11
	i             // 12
)

// Size is our type for 8 bit length unsigned integers.
type Size uint8

// Enum in Golang: IOTA provides an automated way to create a enum in Golang.
const (
	small Size = iota
	medium
	large
	extraLarge
)

func main() {
	fmt.Println("truth:", truth)
	fmt.Println("lf:", lf)
	fmt.Println("statusOK:", statusOK)
	fmt.Println("pi:", pi)
	fmt.Println("comp:", comp)
	fmt.Println("greeting:", greeting)
	fmt.Println(small)
	fmt.Println(medium)
	fmt.Println(large)
	fmt.Println(extraLarge)
}
