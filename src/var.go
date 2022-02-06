package main

import "fmt"

// [var] Variables

// Declare one by one.
var male bool

// Name is exported if it begins with a Capital Letter.
var Name string

// Declare list of variables.
var (
	age      int
	name     string
	location string
)

// [=] Declare and initialize.
var (
	age2      int    = 32
	name2     string = "Interpreted Literal"
	location2 string = `Raw String Literal`
)

// [=] Declare and initialize inferred.
var (
	age3      = 42
	name3     = "Initialized Literal"
	location3 = "File"
)

// <ARRAY>
var myArr [10]int
var sentence [2]string
var multiDimension [2][3]string

// <SLICE>
var mySlice []int // <- nil slice, mySlice == nil: true

func main() {
	// [:=] Inside a function, the := SHORT VARIABLE DECLARATION or
	// short assignment statement.
	name, location := "Its Me", "Here"
	fmt.Println("name:", name, "location:", location)
	raw := `Raw string\tliterals\n`
	fmt.Println("raw string: ", raw)
	hereDoc := `Here docs
or here strings can be
written easily by the
raw string literal.
`
	fmt.Print("here doc: ", hereDoc)
	interp := "Interpreted string\tliterals\n"
	fmt.Print("interpreted string: ", interp)
	lenRaw := len(raw)
	fmt.Printf("Length of raw string: %v\n", lenRaw)
	firstByte := raw[0] // strings can be handled as byte arrays
	fmt.Println("First byte of raw string:", firstByte)
	age := 32
	fmt.Println("Age:", age)

	// <ARRAY> initialization
	sentence[0] = "Hello"
	sentence[1] = "World"
	myArr = [10]int{75, 42, 53, 67} // <ARRAY> initialization by literal
	lenMyArr := len(myArr)          // 10
	fmt.Printf("Length of my array: %v\n", lenMyArr)
	sen := [2]string{"hello", "world!"}
	fmt.Println("String array of sentence, Println:", sen)
	// [hello world!]
	fmt.Printf("String array of sentence, Printf, %v: %s\n", "%s", sen)
	// [hello world!]
	fmt.Printf("String array of sentence, Printf, %v: %q\n", "%q", sen)
	// ["hello" "world!"]
	a := [...]string{"hello", "world!"} // [...] ellipsis:  implicit length
	fmt.Printf("String array with ellipsis, Printf, %v: %q\n", "%q", a)
	multiDimension[0][1] = "42" // reference first row and second column

	// <SLICE> initialization
	mySlice = make([]int, 50, 100)
	mySlice = new([100]int)[0:50]       // equivalent
	mySlice = []int{2, 3, 5, 7, 11, 13} // or initialization by literal
	lenMySlice := len(mySlice)          // 6
	fmt.Printf("Length of my slice: %v\n", lenMySlice)
	fmt.Println("mySlice: ", mySlice)
	// [:] The slice maker operator: array[lowerIndex:higherIndex]
	mySlice = myArr[1:3]         // [:] s[lo:hi]: slice of elements lo - hi-1
	mySlice = append(mySlice, 3) // add the number 3 to the right side
	// Slices are like pointers to named or anonym arrays.
	cities := make([]string, 3)
	cities[0] = "Santa Monica"
	cities[1] = "Venice"
	cities[2] = "Los Angeles"
	// Slices make array manipulation dynamic
	cities = append(cities, "San Diego") // add to the right side
	cities = append(cities, "San Diego", "Mountain View")
	fmt.Printf("cities: %q\n", cities)
}
