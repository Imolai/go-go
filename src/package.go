package main

import (
	"fmt"
	"math/rand" // It is defined in the "rand" directory in the "math" directory
	// as "rand.go" and "package rand" inside.
	/*
		~$ ls ~/sdk/go1.17.6/src/math/rand/rand.go
		/home/$USER/sdk/go1.17.6/src/math/rand/rand.go
		~$ grep ^package ~/sdk/go1.17.6/src/math/rand/rand.go
		package rand
	*/
	"time"
)

func main() {
	// In Go, a name is exported if it begins with a capital letter. Like the
	// Seed, Now, Println, Intn.
	rand.Seed(time.Now().UnixNano())
	fmt.Println("Random number between 0 and 10: ", rand.Intn(10))
}
