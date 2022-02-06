package main

import "fmt"

func main() {
	// [if] [else]
	magicNum := 100
	if magicNum == 100 {
		fmt.Println("Japan")
	} else {
		fmt.Println("Canada")
	}

	// [if] [else] [if]
	if magicNum == 50 {
		fmt.Println("Germany")
	} else if magicNum == 100 {
		fmt.Println("Japan")
	} else {
		fmt.Println("Canada")
	}
}
