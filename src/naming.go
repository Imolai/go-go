package main

import "fmt"

const myName = "Gábor"

func sayHello(name string) string {
	return fmt.Sprintf("Hello, %v!", name)
}

// PrintHello prints a greeting for name.
func PrintHello(name string) {
	fmt.Println(sayHello(name))
}

func main() {
	PrintHello(myName)
}
