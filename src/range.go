package main

import "fmt"

func printList(args ...int) {
	for i, arg := range args {
		fmt.Println("Elem", i, arg)
	}
}

func main() {
	// [range]
	printList(42, 57, 63, 71)
	cities := []string{"Barcelona", "Budapest", "Belgrad", "Wien"}
	fmt.Println("cities printed in range:")
	for i, city := range cities {
		// for _, city := range cities { // skip the index
		fmt.Printf("%v) %v\n", i, city)
	}
}
