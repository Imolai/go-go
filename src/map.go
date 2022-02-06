package main

import "fmt"

//var env map[string]string

func main() {
	// [map]
	//env = make(map[string]string)
	//env["PATH"] = "/bin"
	scientists := map[string]int{
		"Nikola Tesla":        166,
		"Thomas Edison":       175,
		"Tivadar Puskás":      178,
		"James Clerk Maxwell": 191,
	}
	fmt.Printf("%#v\n", scientists)
	env := map[string]string{
		"PATH":  "/bin:/usr/bin",
		"SHELL": "/bin/bash",
		"USER":  "nobody",
	}
	key := "USER"
	elem := "gabor"
	// Insert or update an element in map m:
	env[key] = elem
	// Retrieve an element:
	elem = env[key]
	// Delete an element:
	key = "SHELL"
	delete(env, key)
	// Test that a key is present with a two-value assignment:
	elem1, ok := env[key]
	// If key is in m, ok is true. If not, ok is false and elem is the zero value
	// for the map’s element type.
	if ok == false {
		fmt.Printf("map: No key %v found in env:\n", key)
		fmt.Printf("%q\n", env)
	} else {
		fmt.Println(elem1)
	}
	fmt.Println()
}
