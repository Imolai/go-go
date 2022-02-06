package main

import "fmt"

// [func] function definition
// Can have more parameters.
func add(x int, y int) int {
	// [return]
	return x + y // [+]
}

// More parameters from the same type can be listed with same type.
func myMath(x, y float64) {
	fmt.Println("x + y:", x+y)           // [+] sum
	fmt.Println("x - y:", x-y)           // [-] difference
	fmt.Println("x * y:", x*y)           // [*] product
	fmt.Println("x / y:", x/y)           // [/] quotient
	fmt.Println("x % y:", int(x)%int(y)) // [%] remainder
	/*
		& 	bitwise AND                    // [&]
		| 	bitwise OR                     // [|]
		^ 	bitwise XOR                    // [^]
		&^	bit clear (AND NOT)            // [&^]
	*/
}

// Can return more values.
func sort(x, y int) (int, int) {
	if x < y {
		return x, y // [,]
	}
	return y, x
}

/*
	Multiple return values can be named and act just like variables. If the
	result parameters are named, a return statement without arguments returns the
	current values of the results.
*/
func sort2(x, y int) (first, second int) {
	if x < y {
		first, second = x, y // [,]
	} else {
		first, second = y, x
	}
	return
}

// Variadic functions have variadic parameters.
// [...] ...Type: pack operator
func variadicFunc(arg ...int) {
	for i := range arg {
		fmt.Println("Variadic func", i, arg[i])
	}
}

// <POINTER>
func inc(num *int) {
	/*
		By default Go passes arguments by value (copying the arguments), if you
		want to pass the arguments by reference, you need to pass pointers (or use
		a structure using reference values like slices and maps.
	*/
	// Dereference a pointer, use the [*] symbol.
	*num++ // [++]
}
func dec(num *int) {
	*num-- // [--]
}

// <METHOD>

// user struct is like an OBJECT of attributes.
type user struct {
	id                                  int
	age                                 int
	name, firstName, lastName, location string
}

// NewUser as user's constructor-like function.
func NewUser(id, age int, firstName, lastName, location string) *user {
	u := new(user)
	u.id = id
	u.age = age
	u.firstName = firstName
	u.lastName = lastName
	u.location = location
	return u
}

/*
	Stringer [interface] is defined by the fmt package.
	This is one of the most ubiquitous interfaces.
	A STRINGER is a type that can describe itself as a string.
*/
// Stringer interface defines String() method.
type Stringer interface {
	String() string
}

func (u *user) String() string {
	return fmt.Sprintf("--- USER ---\nname:\t\t%v\nid:\t\t%v\nage:\t\t%v\nlocation:\t%v\n------------", u.Name(), u.id, u.age, u.location)
}

// Namer [interface] contains one method definition.
// An interface type is defined by a set of methods.
type Namer interface {
	Name() string
}

// Name <METHOD> prints the user's name.
// Interfaces are satisfied implicitly.
func (u *user) Name() string {
	return fmt.Sprintf("%s %s", u.firstName, u.lastName)
}

// Greetings <METHOD> is just a function with a RECEIVER ARGUMENT.
func (u *user) Greetings() string { // (u *user): method receiver
	return fmt.Sprintf("Hi %s from %s", u.Name(), u.location)
}

// Birth returns by the birth year.
func (u *user) Birth() int {
	return 2022 - u.age // [-]
}

// <GENERIC>
// Generics will be introduced in Go v1.18.
/*
func printAnySlice[T any](s []T) {
	for _, v := range s {
		fmt.Print(v)
	}
}
*/

func main() {
	// [func] call
	fmt.Println("add(42, 13): ", add(42, 13))
	fmt.Println("mymath(12.0, 33.0): ")
	myMath(12.0, 33.0)
	a, b := sort(42, 13)
	fmt.Println("sort(42, 13): ", a, b)
	a, b = sort2(42, 13)
	fmt.Println("sort2(42, 13): ", a, b)
	fmt.Println("variadicfunc(1, 2, 3):")

	variadicFunc(1, 2, 3)
	var nums []int = []int{42, 57, 93}
	fmt.Println("variadicFunc(42, 57, 93):")
	variadicFunc(nums...) // [...] Var...: unpack operator

	age := 33
	fmt.Println("age: ", age)
	fmt.Println("inc(&age)")
	inc(&age) // get the <POINTER> of a value, use the & symbol
	fmt.Println("age: ", age)
	fmt.Println("dec(&age)")
	dec(&age) // get the <POINTER> of a value, use the & symbol
	fmt.Println("age: ", age)

	// Methods
	/*
		Instead of:
		me := &user{} //me := new(user)
		me.id = 1
		me.age = 127
		me.firstName = "Gábor"
		me.lastName = "Imolai"
		me.location = "Hungary"
	*/
	// We call "constructor".
	me := NewUser(1, 127, "Gábor", "Imolai", "Hungary")
	fmt.Println(me) // Print via Stringer.
}
