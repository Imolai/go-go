package main

import "fmt"

func printNumsCh(from, to int, ch chan bool) {
	for i := from; i <= to; i++ {
		fmt.Print(i, " ")
	}
	ch <- true // [<-]
}

func main() {
	// [chan]
	/*
		To enable communication between Goroutines, Go provides Channels. A channel
		is like a pipe, allowing you to send and receive data from Goroutines.
	*/
	fmt.Println("Print nums by Goroutines using Channels.")
	ch := make(chan bool)
	workers := 2
	go printNumsCh(0, 5, ch)
	go printNumsCh(6, 10, ch)
	// [<-]
	for i := workers; i > 0; i-- {
		<-ch // or: value := <-ch
	}
	fmt.Println()
}
