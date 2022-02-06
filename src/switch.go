package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	// [switch] [case]
	score := 7
	switch score {
	case 1, 3, 5, 7:
		fmt.Println("Terrible")
	}

	// Switch with a short statement.
	switch os := runtime.GOOS; os {
	case "linux":
		fmt.Println("Linux.")
	}

	t := time.Now()
	// Switch without a condition is the same as switch true.
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	}
}
