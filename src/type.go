package main

import (
	"fmt"
	"strconv"
)

// [type]
type myFloat float64
type text []string

var (
	debug    bool = true
	runCount int
	version  float32   = 1.0
	unreal   complex64 = 1.e+1i
	lf       rune      = '\n'
	name     string    = "Gábor"
)

func main() {
	var pi myFloat = 3.14
	var greeting text = text{"Welcome", "to", "the", "Go"}
	fmt.Printf("pi: %v\n", pi)
	fmt.Printf("greeting: %q\n", greeting)

	// type casting
	a := 42
	f := float64(a)
	u := uint(f)
	fmt.Println(a, f, u)

	// type conversion
	i, err := strconv.Atoi("-42")
	fmt.Println(i, err)
	s := strconv.Itoa(-42)
	fmt.Println(s)
	b, err := strconv.ParseBool("true")
	fmt.Println(b, err)
	f, err = strconv.ParseFloat("3.1415", 64)
	fmt.Println(f, err)
	j, err := strconv.ParseInt("-42", 10, 64)
	fmt.Println(j, err)
	ui, err := strconv.ParseUint("42", 10, 64)
	fmt.Println(ui, err)
	s = strconv.FormatBool(true)
	fmt.Println(s)
	s = strconv.FormatFloat(3.1415, 'E', -1, 64)
	fmt.Println(s)
	s = strconv.FormatInt(-42, 16)
	fmt.Println(s)
	s = strconv.FormatUint(42, 16)
	fmt.Println(s)

	// type assertion
	var ifc interface{} = "hello" // static type: interface, dynamic type: string
	s = ifc.(string)
	fmt.Println(s)
	s, ok := ifc.(string)
	fmt.Println(s, ok)
	f, ok = ifc.(float64)
	fmt.Println(f, ok)
	//f = ifc.(float64) // panic!
	//fmt.Println(f)
}
