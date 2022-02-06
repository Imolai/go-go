package main

import "fmt"

// [goto]
func funGoTo() {
	fmt.Println("Before a goto.")
	goto FINISH
	fmt.Println("Unreachable code due to goto.")
FINISH: // label
	fmt.Println("After a goto.")
}

func main() {
	funGoTo()
}
