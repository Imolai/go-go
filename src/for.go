package main

func main() {
	// [for]
	sum := 0
	// Traditional for loop
	for i := 0; i < 10; i++ {
		sum += i
	}
	sum = 1
	// Loop without pre/post statements
	for sum < 1000 {
		sum += sum
	}
	sum = 1
	// For loop as a while loop
	for sum < 1000 {
		sum += sum
	}
	// Infinite loop
	//for {
	// do something in a loop forever
	//}
}
