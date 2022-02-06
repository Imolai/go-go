package main

import (
	"fmt"
	"os"
)

// [defer]
/*
	Defer statement is used to execute a function call just before the
	surrounding function where the defer statement is present returns
	When a function has multiple defer calls, they are pushed on to a stack an
	executed in Last In First Out (LIFO) order.
*/
func writeToTempFile(text string) error {
	file, err := os.Open("temp.txt")
	if err != nil {
		return err
	}
	defer file.Close()

	n, err := file.WriteString(text)
	if err != nil {
		return err
	}
	fmt.Printf("Number of bytes written: %d", n)
	return nil
}

func main() {
	writeToTempFile("Some text from deferred function.")
}
