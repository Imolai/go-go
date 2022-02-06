package main

import (
	"fmt"
	"time"
)

// [go]
func printNums(from, to int) {
	for i := from; i <= to; i++ {
		time.Sleep(50 * time.Millisecond)
		fmt.Print(i, " ")
	}
}

func main() {
	fmt.Println("Print nums by Goroutines.")
	go printNums(0, 5)
	go printNums(6, 10)
	// No output, because the main() does not wait, exits before the Goroutines.
	// It is a simple solution to put a wait in the main thread:
	time.Sleep(500 * time.Millisecond)

	// Anonymous Goroutine
	go func() {
		for i := 11; i < 16; i++ {
			fmt.Print(i, " ")
			time.Sleep(100 * time.Millisecond)
		}
	}()
	time.Sleep(1 * time.Second)
}
