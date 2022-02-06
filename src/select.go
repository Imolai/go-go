package main

import "fmt"

func sumEven(from, to int, ch chan int) {
	sum := 0
	for i := from; i <= to; i++ {
		if i%2 == 0 {
			sum += i
		}
	}
	ch <- sum // [<-]
}

func sumOdd(from, to int, ch chan int) {
	sum := 0
	for i := from; i <= to; i++ {
		if i%2 != 0 {
			sum += i
		}
	}
	ch <- sum // [<-]
}

func main() {
	// [select]
	/*
		The select statement is used to wait on multiple channel operations.
		The syntax is similar to switch except that each of the case statements
		will be a channel operation.
	*/
	fmt.Println("Print even and odd sums by Goroutines using Channels.")
	even := make(chan int)
	odd := make(chan int)
	workers := 2
	go sumEven(0, 100, even)
	go sumOdd(0, 100, odd)
	for i := workers; i > 0; i-- {
		select {
		case x := <-even:
			fmt.Println("Even:", x)
		case y := <-odd:
			fmt.Println("Odd:", y)
		}
	}
}
