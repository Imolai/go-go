package main

import "fmt"

func main() {
	// [break]
	for i := 1; i <= 5; i++ {
		// terminates the loop when i is equal to 3
		if i == 3 {
			break
		}
		fmt.Println(i)
	}
}
