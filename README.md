---
permalink: /index.html
---

<h1 align="center">Go, Go!</h1>

<p align="center"><i>Short basic introduction to Go v1.18 programming language</i></p>

<p align="center"><img src="https://upload.wikimedia.org/wikipedia/commons/thumb/0/05/Go_Logo_Blue.svg/678px-Go_Logo_Blue.svg.png" alt="Go"></p>
<p align="center"><i>Go Logo is Copyright 2018 The Go Authors. All rights reserved.</i></p>

## About

This work was written because it is said everywhere that Go is small and quick
to learn, but I could not find any introduction that reflected this for me. I
have therefore collected and reorganized the most essential basic information,
somewhat as it is presented in [The Specification](https://go.dev/ref/spec)
itself, similar to a Tutorial, but much less verbose and convoluted, somewhere
between the Examples and Specifications genre. I recommend trying out all the
examples on [Playground](https://go.dev/play/), or even better, in your own
development environment, to better absorb the knowledge. Many examples or good
descriptive passages are taken from the great
[referenced works](#used-and-proposed-sources) below, as I could not think of a
better one, nor did I want to surpass them.

[![License](http://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)

<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
## Contents

- [Comments](#comments)
- [Literals](#literals)
- [Identifiers](#identifiers)
  - [Naming](#naming)
  - [Operators](#operators)
    - [Assignment](#assignment)
    - [Arithmetical](#arithmetical)
    - [Comparison](#comparison)
    - [Logical](#logical)
    - [Syntactical](#syntactical)
    - [Pointer](#pointer)
    - [Channel](#channel)
    - [Operator precedence](#operator-precedence)
    - [Blank identifier](#blank-identifier)
  - [Constants](#constants)
  - [Variables](#variables)
    - [Initialization to zero value](#initialization-to-zero-value)
  - [Types](#types)
    - [Type Conversion - Casting](#type-conversion---casting)
    - [Type Conversion - strconv](#type-conversion---strconv)
    - [Type Assertions](#type-assertions)
  - [Structs](#structs)
    - [Named Structs](#named-structs)
    - [Anonymous Structs](#anonymous-structs)
  - [Maps](#maps)
  - [Functions / Methods](#functions--methods)
    - [Named Functions](#named-functions)
    - [Anonymous Functions](#anonymous-functions)
    - [Goto in a Function](#goto-in-a-function)
    - [Return a Function or Method](#return-a-function-or-method)
    - [Defer a Function or Method](#defer-a-function-or-method)
  - [Interfaces](#interfaces)
- [If Conditional Execution](#if-conditional-execution)
  - [Else of If](#else-of-if)
- [Switch Multi-way Execution](#switch-multi-way-execution)
  - [Expression Switches](#expression-switches)
  - [Type Switches](#type-switches)
  - [Case in a Switch](#case-in-a-switch)
  - [Default case in a Switch](#default-case-in-a-switch)
  - [Fallthrough in a Switch](#fallthrough-in-a-switch)
- [For The Loop](#for-the-loop)
  - [Break a Loop](#break-a-loop)
  - [Continue a Loop](#continue-a-loop)
  - [Ranges](#ranges)
- [Packages](#packages)
  - [Import paths](#import-paths)
  - [Exported names](#exported-names)
  - [Modules](#modules)
  - [Creating a module](#creating-a-module)
- [Import Packages](#import-packages)
- [Goroutines](#goroutines)
  - [Channels](#channels)
    - [Deadlock](#deadlock)
  - [Select of Communications](#select-of-communications)
- [Used and proposed sources](#used-and-proposed-sources)
- [Author](#author)
- [Support](#support)
- [License](#license)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

## Comments

Comments serve as program documentation.

---

[comment.go](./src/comment.go)

```go
// Package main provides possibility to see Go documentation.
package main

import "fmt"

// Version represents the version number of the package.
const Version = 1.0

// Help returns nothing, but prints the usage as side effect.
func Help() {
	fmt.Println("My main package. Created by me.")
	fmt.Printf("Version: %.1f\n", Version)
}

func main() {
	/*
		General comments start with the character sequence slash-star and stop with
		the first subsequent character sequence star-slash.
		To document a type, variable, constant, function, or even a package, write
		a regular comment directly preceding its declaration, with no intervening
		blank line.
	*/
	// Line comments start with the character sequence slash-slash and stop at
	// the end of the line.
	Help()
}
```

**[⬆ Back to top](#contents)**

## Literals

A literal is a fixed value assigned to variables or constants.

---

[literal.go](./src/literal.go)

```go
package main

import "fmt"

func main() {
	// Integer literals
	var i uint64
	i = 42
	i = 4_2
	i = 0600
	i = 0_600
	i = 0o600
	i = 0O600 // second character is capital letter 'O'
	i = 0xBadFace
	i = 0xBad_Face
	i = 0x_67_7a_2f_cc_40_c6
	i = 17014118346046923173
	i = 170_141183_460469_23173
	fmt.Println(i)

	// Floating-point literals
	var f float64
	f = 0.
	f = 72.40
	f = 072.40 // == 72.40
	f = 2.71828
	f = 1.e+0
	f = 6.67428e-11
	f = 1e6
	f = .25
	f = .12345e+5
	f = 1_5.        // == 15.0
	f = 0.15e+0_2   // == 15.0
	f = 0x1p-2      // == 0.25
	f = 0x2.p10     // == 2048.0
	f = 0x1.Fp+0    // == 1.9375
	f = 0x.8p-0     // == 0.5
	f = 0x15e - 2   // == 0x15e - 2 (integer subtraction)
	f = 0x_1FFFp-16 // == 0.1249847412109375
	fmt.Println(f)

	// Imaginary literals
	var c complex128
	c = 0i
	c = 123i   // == 123i for backward-compatibility
	c = 0o123i // == 0o123 * 1i == 83i
	c = 0xabci // == 0xabc * 1i == 2748i
	c = 0.i
	c = 2.71828i
	c = 1.e+0i
	c = 6.67428e-11i
	c = 1e6i
	c = .25i
	c = .12345e+5i
	c = 0x1p-2i // == 0x1p-2 * 1i == 0.25i
	fmt.Println(c)

	// Rune literals
	var r rune
	r = 'a'
	r = '本'
	r = '\'' // rune literal containing single quote character
	r = '\t'
	r = '\000'
	r = '\007'
	r = '\377'
	r = '\x07'
	r = '\xff'
	r = '\u12e4'
	r = '\U00101234'
	r = 'ä'
	fmt.Println(r)

	// String literals
	var s string
	s = `abc` // same as "abc"
	s = `\n
\n` // same as "\\n\n\\n"
	s = "\n"
	s = "\"" // same as `"`
	s = "Hello, world!\n"
	s = "日本語"
	s = "\u65e5本\U00008a9e"
	s = "\xff\u00FF"
	s = "日本語"                               // UTF-8 input text
	s = `日本語`                               // UTF-8 input text as raw literal
	s = "\u65e5\u672c\u8a9e"                   // the explicit Unicode code points
	s = "\U000065e5\U0000672c\U00008a9e"       // the explicit Unicode code points
	s = "\xe6\x97\xa5\xe6\x9c\xac\xe8\xaa\x9e" // the explicit UTF-8 bytes
	fmt.Println(s)
}
```

**[⬆ Back to top](#contents)**

## Identifiers

### Naming

When a package is imported, only entities (functions, types, variables,
constants) whose names start with a capital letter can be used / accessed. The
recommended style of naming in Go is that identifiers will be named using
camelCase, except for those meant to be accessible across packages which should
be PascalCase.

**Module Name**: Name your project in lowercase, otherwise it would look really
weird in the imports.

Example: `golang.org/x/tools`

**Package Name**: For package, use singular noun instead of plural. Use user
instead of users, notification instead of notifications.

**File Name**: For file name, try to keep it short, but still meaningful.
Use lower case and use snake_case instead of mixedCaps. So, use server.go,
notification.go or notification_server.go instead of notificationServer.go.

**Function/Method**: Use mixedCaps for naming your function. Use getArticles or
GetArticles instead of get_articles or get-articles.

**Variable**: Name of variable should be short but should be meaningful enough.
So, you can use srv or svc instead of service, server or myService. User or repo
instead of repository or myRepository.

**Constant**: Constant follows the same rule with variables and functions, hence
use mixedCaps for naming. No underscore unless it's generated code. So, use
StatusActive instead of STATUS_ACTIVE.

**Error**: Name of error variable should start with Err. Use ErrNotFound instead
of NotFoundError. Note that it's **Err**, not *Error*.

**Interfaces**: The simplest rule for naming interface is adding *'-er'* into
the name of the action/method it represents.

See: <https://pthethanh.herokuapp.com/blog/articles/golang-name-conventions>

---

[naming.go](./src/naming.go)

```go
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
```

**[⬆ Back to top](#contents)**

### Operators

Operators combine operands into expressions.

```
Expression = UnaryExpr | Expression binary_op Expression .
UnaryExpr  = PrimaryExpr | unary_op UnaryExpr .

binary_op  = "||" | "&&" | rel_op | add_op | mul_op .
rel_op     = "==" | "!=" | "<" | "<=" | ">" | ">=" .
add_op     = "+" | "-" | "|" | "^" .
mul_op     = "*" | "/" | "%" | "<<" | ">>" | "&" | "&^" .

unary_op   = "+" | "-" | "!" | "^" | "*" | "&" | "<-" .
```

#### Assignment

|Operator|Description|
|--------|-----------|
|`=`|normal assignment|
|`:=`|short assignment (short variable declaration, inside only functions)|
|`+=`|addition assignment|
|`-=`|subtraction assignment|
|`*=`|multiplication assignment|
|`/=`|division assignment|
|`%=`|remainder assignment|
|`&=`|bitwise AND assignment|
|`\|=`|bitwise OR assignment|
|`^=`|bitwise XOR assignment|
|`<<=`|left shift assignment|
|`>>=`|right shift assignment|
|`&^=`|bit clear (AND NOT) assignment|
|`++`|increase by one| *form statement, not operator*
|`--`|decrease by one| *form statement, not operator*

**[⬆ Back to top](#contents)**

#### Arithmetical

|Operator|Description|
|--------|-----------|
|`+`|addition|
|`-`|subtraction|
|`*`|multiplication|
|`/`|division|
|`%`|remainder|
|`&`|bitwise AND|
|`\|`|bitwise OR|
|`^`|bitwise XOR|
|`<<`|left shift|
|`>>`|right shift|
|`&^`|bit clear (AND NOT)|

**[⬆ Back to top](#contents)**

#### Comparison

|Operator|Description|
|--------|-----------|
|`==`|equal|
|`!=`|not equal|
|`<`|less than|
|`>`|greater than|
|`<=`|less than or equal|
|`>=`|greater than or equal|

**[⬆ Back to top](#contents)**

#### Logical

|Operator|Description|
|--------|-----------|
|`&&`|logical and|
|`\|\|`|logical or|
|`!`|logical not|

**[⬆ Back to top](#contents)**

#### Syntactical

|Operator|Description|
|--------|-----------|
|`...`|(ellipsis) pack/unpack parameters, arguments; implicit length|
|`,`|(comma) identifier list separator|
|`.`|(dot) import; struct-field, struct-method, qualified identifier separator|
|`;`|(semicolon) statement separator|
|`:`|(colon) slice maker; label end|
|`(`|parameter/argument list separator|
|`)`|parameter/argument list separator|
|`[`|array, slice declaration, indexing; generic type|
|`]`|array, slice declaration, indexing; generic type|
|`{`|block separator|
|`}`|block separator|

**[⬆ Back to top](#contents)**

#### Pointer

|Operator|Description|
|--------|-----------|
|`&`|address of / create pointer|
|`*`|dereference pointer|

**[⬆ Back to top](#contents)**

#### Channel

|Operator|Description|
|--------|-----------|
|`<-`|send / receive operator|

#### Operator precedence

**Unary operators** have the **highest** precedence. As the ++ and -- operators form
statements, not expressions, they fall outside the operator hierarchy. As a
consequence, statement \*p++ is the same as (\*p)++.

There are five precedence levels for binary operators. Multiplication operators
bind strongest, followed by addition operators, comparison operators,
&& (logical AND), and finally || (logical OR):

```
Precedence    Operator
    5             *  /  %  <<  >>  &  &^
    4             +  -  |  ^
    3             ==  !=  <  <=  >  >=
    2             &&
    1             ||
```

**[⬆ Back to top](#contents)**

#### Blank identifier

The blank identifier is represented by the underscore character `_`. It serves
as an anonymous placeholder instead of a regular (non-blank) identifier.

**[⬆ Back to top](#contents)**

### Constants

> KEYWORDS: **const**.

There are boolean constants, rune constants, integer constants, floating-point
constants, complex constants, and string constants. Rune, integer,
floating-point, and complex constants are collectively called numeric constants.

The **predeclared** constants are: `true`, `false`, `iota`.

The special constant `iota`:

- A counter which starts with zero.
- Increases by 1 after each line.
- Is only used with constant.

Iota keyword can be used on each line, or can be skipped as well.

---

[const.go](./src/const.go)

```go
package main

import "fmt"

// [const] Constants

// Pi exported constant, one value per line.
const Pi = 3.14 // const cannot be declared using the := syntax

// More constants by one const.
const (
	truth    = false           // boolean constant
	lf       = '\n'            // rune constant
	statusOK = 200             // integer constant
	pi       = 3.1415965359    // floating-point constant
	comp     = 1.e+0i          // complex constant
	greeting = "Hello, world!" // string constant
)

const (
	a = 0
	b = 1
	c = 2
)

// The same with auto increment IOTA.
const (
	d = iota // 0
	e        // 1
	f        // 2
)

// Iota can also start from non-zero number - iota expressions can also be used
// to start iota from any number
const (
	g = iota + 10 // 10
	h             // 11
	i             // 12
)

// Size is our type for 8 bit length unsigned integers.
type Size uint8

// Enum in Golang: IOTA provides an automated way to create a enum in Golang.
const (
	small Size = iota
	medium
	large
	extraLarge
)

func main() {
	fmt.Println("truth:", truth)
	fmt.Println("lf:", lf)
	fmt.Println("statusOK:", statusOK)
	fmt.Println("pi:", pi)
	fmt.Println("comp:", comp)
	fmt.Println("greeting:", greeting)
	fmt.Println(small)
	fmt.Println(medium)
	fmt.Println(large)
	fmt.Println(extraLarge)
}
```

**[⬆ Back to top](#contents)**

### Variables

> KEYWORDS: **var**.

A variable is a storage location for holding a value. The set of permissible
values is determined by the variable's type.

The **predeclared** zero value: `nil`.

#### Initialization to zero value

When storage is allocated for a variable, and no explicit initialization is
provided, the variable or value is given a default value. Each element of such
a variable or value is *set to the **zero** value* for its type: `false` for
`bool`eans, `0` for *numeric* types, `""` for `string`s, and `nil` for
*pointer*s, `func`tions, `interface`s, *slice*s, `chan`nels, and `map`s. This
initialization is done recursively.

---

[var.go](./src/var.go)

```go
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
```

**[⬆ Back to top](#contents)**

### Types

> KEYWORDS: **type**.

A type determines a set of values together with operations and methods specific
to those values. A type may be denoted by a type name, if it has one, or
specified using a type literal, which composes a type from existing types.

**Basic types**: boolean, int, float, complex, rune, string.  
**Composite types**: array, struct, pointer, function, interface, slice, map,
channel types.  
**Static type** (or just *type*) of a variable is the type given in its
declaration.  
**Dynamic type**, which is the concrete type of the value assigned to the
variable at run time.

The *predeclared* types:

- bool,
- byte, rune,
- float32, float64,
- int, int8, int16, int32, int64,
- uint, uint8, uint16, uint32, uint64, uintptr,
- complex64, complex128,
- string,
- error.

The documentation of builtin types: <https://pkg.go.dev/builtin>.

uint8    the set of all unsigned  8-bit integers (0 to 255)  
uint16   the set of all unsigned 16-bit integers (0 to 65535)  
uint32   the set of all unsigned 32-bit integers (0 to 4294967295)  
uint64   the set of all unsigned 64-bit integers (0 to 18446744073709551615)  

int8     the set of all signed  8-bit integers (-128 to 127)  
int16    the set of all signed 16-bit integers (-32768 to 32767)  
int32    the set of all signed 32-bit integers (-2147483648 to 2147483647)  
int64    the set of all signed 64-bit integers (-9223372036854775808 to
         9223372036854775807)

float32  the set of all IEEE-754 32-bit floating-point numbers  
float64  the set of all IEEE-754 64-bit floating-point numbers  

complex64   the set of all complex numbers with float32 real and imaginary parts
complex128  the set of all complex numbers with float64 real and imaginary parts

byte     alias for uint8
rune     alias for int32  

uint     either 32 or 64 bits  
int      same size as uint
uintptr  an unsigned integer large enough to store the uninterpreted bits of a
         pointer value

**[⬆ Back to top](#contents)**

#### Type Conversion - Casting

The expression `T(v)` converts the value `v` to the type `T`.

Some numeric conversions:

```go
var i int = 42
var f float64 = float64(i)
var u uint = uint(f)
```

Or, with short variable declaration:

```go
i := 42
f := float64(i)
u := uint(f)
```

The assignment between items of different type requires an explicit conversion.

**[⬆ Back to top](#contents)**

#### Type Conversion - strconv

Package [strconv](https://pkg.go.dev/strconv) implements conversions to and from
string representations of basic data types.

```go
i, err := strconv.Atoi("-42")
s := strconv.Itoa(-42)
b, err := strconv.ParseBool("true")
f, err := strconv.ParseFloat("3.1415", 64)
i, err := strconv.ParseInt("-42", 10, 64)
u, err := strconv.ParseUint("42", 10, 64)
s := strconv.FormatBool(true)
s := strconv.FormatFloat(3.1415, 'E', -1, 64)
s := strconv.FormatInt(-42, 16)
s := strconv.FormatUint(42, 16)
q := strconv.Quote("Hello, 世界")
q := strconv.QuoteToASCII("Hello, 世界")
```

**[⬆ Back to top](#contents)**

#### Type Assertions

A type assertion provides access to an interface value's underlying concrete
value.

```go
t := i.(T)
```

This statement asserts that the interface value `i` holds the concrete type `T`
and assigns the underlying `T` value to the variable `t`.

If `i` does not hold a `T`, the statement will trigger a panic.

To test whether an interface value holds a specific type, a type assertion can
*return two values*: the *underlying value* and a *boolean value* that reports
whether the assertion succeeded.

```go
t, ok := i.(T)
```

If `i` holds a `T`, then `t` will be the underlying value and ok will be true.

If not, ok will be `false` and `t` will be the zero value of type `T`, and no
panic occurs.

> Note the similarity between this syntax and that of reading from a map.

```go
m := make(map[string]int) // map
m["Answer"] = 42          // map has one entry
v, ok := m["Answer"]      // reading from map similar
```

---

[type.go](./src/type.go)

```go
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
```

**[⬆ Back to top](#contents)**

### Structs

> KEYWORDS: **struct**.

A struct is a sequence of named elements, called *fields*, each of which has a
name and a type.

#### Named Structs

---

[struct.go](./src/struct.go)

```go
package main

import "fmt"

// [struct]
var aPoint struct {
	x int
	y int
}

type point struct {
	x int
	y int
}

// User struct is like an object of attributes.
type User struct {
	ID                                  int
	Age                                 int
	Name, FirstName, LastName, Location string
}

// Player struct includes User and additional attribute
type Player struct {
	User
	GameID int
}

func main() {
	// [struct] initialization
	p := point{1, 2}
	q := &point{1, 2} // has type *point
	r := point{x: 1}  // y:0 is implicit
	s := point{}      // x:0 and y:0
	fmt.Printf("point{1, 2}: %v,\n", p)
	fmt.Printf("&point{1, 2}: %v,\n", q)
	fmt.Printf("point{x: 1}: %v,\n", r)
	fmt.Printf("point{}: %v\n", s)
	me := Player{}
	me.ID = 47
	me.Name = "Gábor"
	me.Location = "Hungary"
	fmt.Println(me)
}
```

**[⬆ Back to top](#contents)**

#### Anonymous Structs

---

[struct.go](./src/struct_anon.go)

```go
package main

import "fmt"

// --- Anonymous Field in Struct ---

type Polygon struct {
	Sides int
}

func (p *Polygon) NSides() int {
	return p.Sides
}

type Triangle struct {
	Polygon // Anonymous Field
}

// ---------------------------------

func main() {
	pizza := struct {
		name string
	}{
		name: "Pizza",
	}
	fmt.Println(pizza) // prints {Pizza}

	// Using Anonymous Field of Struct to reduce unnecessary function calls.
	t := Triangle{
		Polygon{
			Sides: 3,
		},
	}
	fmt.Println(t.NSides()) // 3
}
```

**[⬆ Back to top](#contents)**

### Maps

> KEYWORDS: **map**.

A map is an unordered group of *elements* of one type, called the element type,
indexed by a set of unique *keys* of another type, called the key type.

---

[map.go](./src/map.go)

```go
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
```

**[⬆ Back to top](#contents)**

### Functions / Methods

> KEYWORDS: **func**.

A function type denotes the set of all functions with the same parameter and
result types.

The **predeclared** functions: append, cap, close, complex, copy, delete, imag,
len, make, new, panic, print, println, real, recover.

The documentation of builtin functions: <https://pkg.go.dev/builtin>.

- `append(slice []Type, elems ...Type) []Type`	- The append built-in function
  appends elements to the end of a slice.
- `cap(v Type) int`	- The cap built-in function returns the capacity of v,
  according to its type
- `close(c chan<- Type)`	- The close built-in function closes a channel, which
  must be either bidirectional or send-only.
- `complex(r, i FloatType) ComplexType`	- The complex built-in function
  constructs a complex value from two floating-point values.
- `copy(dst, src []Type) int`	- The copy built-in function copies elements from
  a source slice into a destination slice.
- `delete(m map[Type]Type1, key Type)`	- The delete built-in function deletes
  the element with the specified key (m[key]) from the map.
- `imag(c ComplexType) FloatType`	- The imag built-in function returns the
  imaginary part of the complex number c.
- `len(v Type) int`	- The len built-in function returns the length of v,
  according to its type. In the most cases, this means the NUMBER OF ELEMENTS,
  except of STRINGs, where it is the NUMBER OF BYTES.
- `make(t Type, size ...IntegerType) Type`	- The make built-in function
  allocates and initializes an object of type slice, map, or chan (only).
- `new(Type) *Type`	- The new built-in function allocates memory.
- `panic(v interface{})`	- The panic built-in function stops normal execution
  of the current goroutine.
- `print(args ...Type)`	- The print built-in function formats its arguments in
  an implementation-specific way and writes the result to standard error.
- `println(args ...Type)`	- The println built-in function formats its arguments
  in an implementation-specific way and writes the result to standard error.
- `real(c ComplexType) FloatType`	- The real built-in function returns the real
  part of the complex number c.
- `recover() interface{}`	- The recover built-in function allows a program to
  manage behavior of a panicking goroutine.

**[⬆ Back to top](#contents)**

#### Named Functions

---

[func.go](./src/func.go)

```go
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
// Generics introduced in Go v1.18.

func printAnySlice[T any](s []T) {
	for _, v := range s {
		fmt.Print(v, " ")
	}
	fmt.Println("")
}

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
	printAnySlice(nums)
}
```

**[⬆ Back to top](#contents)**

#### Anonymous Functions

---

[func_anon.go](./src/func_anon.go)

```go
package main

import "fmt"

func main() {
	// Anonymous function
	v := func() {
		fmt.Println("Inside anonymous function")
	}
	v() // call
	func() {
		fmt.Println("Anonymous function gets invoked immediately")
	}()
	func(v int) {
		fmt.Println(v)
	}(42) // Passing Arguments to an Anonymous Function
	g := func(v string) {
		fmt.Println(v)
	}
	func(v string, g func(v string)) {
		g(v)
	}("Passing Anonymous Functions as an Argument", g)
	f1 := func() func(v string) {
		f := func(v string) {
			fmt.Println(v)
		}
		return f
	}
	h := f1()
	h("Returning Anonymous Function from a Function")
}
```

**[⬆ Back to top](#contents)**

#### Goto in a Function

> KEYWORDS: **goto**.

A "goto" statement transfers control to the statement with the corresponding
label within the same function.

---

[goto.go](./src/goto.go)

```go
package main

import "fmt"

// [goto]
func funGoTo() {
	fmt.Println("Before a goto.")
	goto FINISH
	fmt.Println("Unreachable code due to goto.")
FINISH: // label
	fmt.Println("After a goto.")
}

func main() {
	funGoTo()
}
```

**[⬆ Back to top](#contents)**

#### Return a Function or Method

> KEYWORDS: **return**.

A "return" statement in a function F terminates the execution of F, and
optionally provides one or more result values.

---

[return.go](./src/return.go)

```go
package main

import "fmt"

func add(x int, y int) int {
	// [return]
	return x + y // [+]
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

func main() {
	fmt.Println("add(42, 13): ", add(42, 13))
	fmt.Println("mymath(12.0, 33.0): ")
	a, b := sort(42, 13)
	fmt.Println("sort(42, 13): ", a, b)
	a, b = sort2(42, 13)
	fmt.Println("sort2(42, 13): ", a, b)
}
```

**[⬆ Back to top](#contents)**

#### Defer a Function or Method

> KEYWORDS: **defer**.

A "defer" statement invokes a function whose execution is deferred to the
moment the surrounding function *returns*, either because the surrounding
function *executed a return statement*, *reached the end* of its function body,
or because the corresponding goroutine is *panicking*.

---

[defer.go](./src/defer.go)

```go
package main

import (
	"fmt"
	"os"
)

// [defer]
/*
	Defer statement is used to execute a function call just before the
	surrounding function where the defer statement is present returns
	When a function has multiple defer calls, they are pushed on to a stack an
	executed in Last In First Out (LIFO) order.
*/
func writeToTempFile(text string) error {
	file, err := os.Open("temp.txt")
	if err != nil {
		return err
	}
	defer file.Close()

	n, err := file.WriteString(text)
	if err != nil {
		return err
	}
	fmt.Printf("Number of bytes written: %d", n)
	return nil
}

func main() {
	writeToTempFile("Some text from deferred function.")
}
```

**[⬆ Back to top](#contents)**

### Interfaces

> KEYWORDS: **interface**.

An interface type specifies a *method set* called its interface. A variable of
interface type can store a value of any type with a method set that is any
superset of the interface. Such a type is said to *implement the interface*.

---

[interface.go](./src/interface.go)

```go
package main

import (
	"fmt"
	"time"
)

// user struct is like an OBJECT of attributes.
type user struct {
	id                                  int
	age                                 int
	name, firstName, lastName, location string
}

// [interface]
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

// The ERROR type is a built-in interface similar to fmt.Stringer:
type error interface {
	Error() string
}
type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}
func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

func main() {
	me := new(user)
	me.id = 1
	me.age = 127
	me.firstName = "Gábor"
	me.lastName = "Imolai"
	me.location = "Hungary"
	fmt.Println(me) // Print via Stringer.
}
```

**[⬆ Back to top](#contents)**

## If Conditional Execution

> KEYWORDS: **if**.

"If" statements specify the conditional execution of two branches according to
the value of a boolean expression.

---

[if.go](./src/if.go)

```go
package main

import "fmt"

func main() {
	// [if]
	answer := 42
	if answer == 42 { // [==]
		fmt.Println("The Answer to the Ultimate Question of Life, the Universe, and Everything...")
	}
	foo := func() interface{} {
		return nil
	}
	// Like for, the if statement can start with a short statement.
	if err := foo(); err != nil { // [!=]
		panic(err)
	}
	// Nested ifs can be avoided by chaining conditions logically.
	question := "What do you get if you multiply six by nine?"
	if question == "What do you get if you multiply six by nine?" && answer == 42 {
		fmt.Println("Don't panic!")
	}
}
```

**[⬆ Back to top](#contents)**

### Else of If

> KEYWORDS: **else**.

If the expression evaluates to true, the "if" branch is executed, otherwise, if
present, the "else" branch is executed.

---

[else.go](./src/else.go)

```go
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
```

**[⬆ Back to top](#contents)**

## Switch Multi-way Execution

> KEYWORDS: **switch**.

"Switch" statements provide multi-way execution. An expression or type is
compared to the "cases" inside the "switch" to determine which branch to
execute.

### Expression Switches

In an expression switch, the cases contain expressions that are compared against
the value of the switch expression.

---

[switch.go](./src/switch.go)

```go
package main

import (
        "fmt"
        "runtime"
        "time"
)

func main() {
        // We can write multiple if statements.
        age := 42
        if age == 16 {
                fmt.Println("Too young")
        }
        if age == 42 {
                fmt.Println("Adult")
        }
        if age == 70 {
                fmt.Println("Senior")
        }

        // But switch is more readable.
        // [switch] [case]
        switch age {
        case 16:
                fmt.Println("Too young")
        case 42:
                fmt.Println("Adult")
        case 70:
                fmt.Println("Senior")
        }

        // Switch with a short statement.
        switch os := runtime.GOOS; os {
        case "linux":
                fmt.Println("Linux.")
        }

        t := time.Now()
        // Switch without a condition is the same as switch true.
        switch {
        case t.Hour() < 12:
                fmt.Println("Good morning!")
        }
}
```

**[⬆ Back to top](#contents)**

### Type Switches

In a type switch, the cases contain types that are compared against the type of
a specially annotated switch expression.

A type switch is a construct that permits several type assertions in series.

A type switch is like a regular switch statement, but the cases in a type
switch specify types (not values), and those values are compared against the
type of the value held by the given interface value.

```go
switch v := i.(type) {
case T:
    // here v has type T
case S:
    // here v has type S
default:
    // no match; here v has the same type as i
}
```

---

[type_switch.go](./src/type_switch.go)

```go
package main

import "fmt"

func do(i interface{}) {
	switch v := i.(type) {
	case bool:
		fmt.Printf("Boolean: %v\n", v)
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

func main() {
	do(21)
	do("hello")
	do(true)
}
```

**[⬆ Back to top](#contents)**

### Case in a Switch

> KEYWORDS: **case**.

In an expression switch, the switch expression is evaluated and the case
expressions, which need not be constants, are evaluated left-to-right and
top-to-bottom; the first one that equals the switch expression triggers
execution of the statements of the associated case; the other cases are skipped.

---

[case.go](./src/case.go)

```go
package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	// [switch] [case]
	score := 7
	switch score {
	case 0, 1, 3:
		fmt.Println("Terrible")
	case 4, 5:
		fmt.Println("Mediocre")
	case 6, 7:
		fmt.Println("Not bad")
	case 8, 9:
		fmt.Println("Almost perfect")
	case 10:
		fmt.Println("hmm did you cheat?")
	}

	// Switch with a short statement.
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	}

	t := time.Now()
	// Switch without a condition is the same as switch true.
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	}
}
```

**[⬆ Back to top](#contents)**

### Default case in a Switch

> KEYWORDS: **default**.

If no case matches and there is a "default" case, its statements are executed.
There can be at most one default case and it may appear anywhere in the "switch"
statement.

---

[default.go](./src/default.go)

```go
package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	// [switch] [case]
	score := 7
	switch score {
	case 0, 1, 3:
		fmt.Println("Terrible")
	case 4, 5:
		fmt.Println("Mediocre")
	case 6, 7:
		fmt.Println("Not bad")
	case 8, 9:
		fmt.Println("Almost perfect")
	case 10:
		fmt.Println("hmm did you cheat?")
	default:
		fmt.Println(score, " off the chart")
	}

	// Switch with a short statement.
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}

	t := time.Now()
	// Switch without a condition is the same as switch true.
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}
```

**[⬆ Back to top](#contents)**

### Fallthrough in a Switch

> KEYWORDS: **fallthrough**.

In a case or default clause, the last non-empty statement may be a (possibly
labeled) "fallthrough" statement to indicate that control should flow from the
end of this clause to the first statement of the next clause. Otherwise control
flows to the end of the "switch" statement. A "fallthrough" statement may appear
as the last statement of all but the last clause of an expression switch.

---

[fallthrough.go](./src/fallthrough.go)

```go
package main

import "fmt"

func main() {
	// [fallthrough]
	n := 4
	switch n {
	case 0:
		fmt.Println("is zero")
		fallthrough
	case 1:
		fmt.Println("is <= 1")
		fallthrough
	case 2:
		fmt.Println("is <= 2")
		fallthrough
	case 3:
		fmt.Println("is <= 3")
		fallthrough
	case 4:
		fmt.Println("is <= 4")
		fallthrough
	case 5:
		fmt.Println("is <= 5")
		fallthrough
	case 6:
		fmt.Println("is <= 6")
		fallthrough
	case 7:
		fmt.Println("is <= 7")
		fallthrough
	case 8:
		fmt.Println("is <= 8")
		fallthrough
	default:
		fmt.Println("Try again!")
	}
}
```

**[⬆ Back to top](#contents)**

## For The Loop

> KEYWORDS: **for**.

A "for" statement specifies repeated execution of a block. There are three
forms: The iteration may be controlled by a single condition, a "for" clause, or
a "range" clause.

---

[for.go](./src/for.go)

```go
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
```

**[⬆ Back to top](#contents)**

### Break a Loop

> KEYWORDS: **break**.

A "break" statement terminates execution of the innermost "for", "switch", or
"select" statement within the same function.

---

[break.go](./src/break.go)

```go
package main

import "fmt"

func main() {
	// [break]
	for i := 1; i <= 5; i++ {
		// terminates the loop when i is equal to 3
		if i == 3 {
			break
		}
		fmt.Println(i)
	}
}
```

**[⬆ Back to top](#contents)**

### Continue a Loop

> KEYWORDS: **continue**.

A "continue" statement begins the next iteration of the innermost "for" loop at
its post statement.

---

[continue.go](./src/continue.go)

```go
package main

import "fmt"

func main() {
	// [continue]
	for i := 1; i <= 5; i++ {
		// skips the iteration when i is equal to 3
		if i == 3 {
			continue
		}
		fmt.Println(i)
	}
}
```

**[⬆ Back to top](#contents)**

### Ranges

> KEYWORDS: **range**.

A "for" statement with a "range" clause iterates through all entries of an
array, slice, string or map, or values received on a channel. For each entry it
assigns iteration values to corresponding iteration variables if present and
then executes the block.

---

[range.go](./src/range.go)

```go
package main

import "fmt"

func printList(args ...int) {
	for i, arg := range args {
		fmt.Println("Elem", i, arg)
	}
}

func main() {
	// [range]
	printList(42, 57, 63, 71)
	cities := []string{"Barcelona", "Budapest", "Belgrad", "Wien"}
	fmt.Println("cities printed in range:")
	for i, city := range cities {
		// for _, city := range cities { // skip the index
		fmt.Printf("%v) %v\n", i, city)
	}
}
```

**[⬆ Back to top](#contents)**

## Packages

> KEYWORDS: **package**.

Go applications are organized in packages. **Package** is a way of grouping
related code.
Every Go source file (\*.go) in a Go application belongs to a package.
Package can be of two types.

- **Executable** package – Only main is the executable package in Go. A .go file
  might belong to the main package present within a specific directory. We
  will see later how the directory name or the .go file name matters. The main
  package will contain a main function that denotes the start of a program. On
  installing the main package it will create an executable in the `$GOBIN`
  directory.
- **Utility** package – Any package other than the main package is a utility
  package. It is not self-executable. It just contains the utility function
  and other utility things that can be utilized by an executable package.

```bash
$ ls ~/sdk/go1.17.6/src/math/*.go | xargs grep ^package | cut -d: -f2 | sort -u
package math
package math_test
```

**[⬆ Back to top](#contents)**

### Import paths

Choose a good import path (make your package "go get"-able). Import paths should
be globally unique, so use the path of your source repository as its base. For
instance, the websocket package from the go.net sub-repository has an import
path of `golang.org/x/net/websocket`. The Go project owns the path
`github.com/golang`, so that path cannot be used by another author for a
different package.

**[⬆ Back to top](#contents)**

### Exported names

Instead of public, private or protected keywords, to control the visibility is
using the capitalized and non-capitalized formats.

- **Capitalized** Identifiers are **exported**. The capital letter indicates
  that this is an exported identifier. It will be visible in other packages.
- **Non-capitalized** identifiers are **not exported**. The lowercase indicates
  that the identifier is not exported and will only be accessed from within the
  same package.

**[⬆ Back to top](#contents)**

### Modules

**Module** is a way of dealing with dependencies.  A module by definition is a
collection of related packages with `go.mod` at its root. The go.mod file
defines the

- Module import path.
- Dependency requirements of the module for a successful build. It defines both
project’s dependencies requirement and also locks them to their correct version

Consider a module as a directory containing a collection of packages. The
packages can be nested as well.
Modules provide

- Dependency Management
- With modules go project doesn’t necessarily have to lie in the `$GOPATH/src`
  folder.

**[⬆ Back to top](#contents)**

### Creating a module

1. Make package directory

   ```bash
   $ mkdir greeting
   $ cd greeting
   ```

2. Initialize module

   ```bash
   greeting$ go mod init example.com/greeting
   ```

3. Make package code(s)

   ```bash
   greeting$ vi greeting.go
   ```

4. Run/build application

   ```bash
   greeting$ go run greeting.go // AND/OR
   greeting$ go build greeting.go
   ```

---

[package.go](./src/package.go)

```go
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
```

**[⬆ Back to top](#contents)**

## Import Packages

> KEYWORDS: **import**.

An import declaration states that the source file containing the declaration
depends on functionality of the imported package and enables access to exported
identifiers of that package.

---

[import.go](./src/import.go)

```go
package main

// [import]
/*
	Almost everything we want to do in Go comes from some package. And to use
	those packages, we need to import them.

	**Aliased import**
	import f "fmt"

	**Blank import**
	import _ "time"    ignore the unused import

	**Dot import** [.]
	If an explicit period (.) appears instead of a name, all the package's
	exported identifiers will be declared in the current file's file block an
	can be accessed without a qualifier
	Import declaration          Local name of Sin
	import   "lib/math"         math.Sin
	import M "lib/math"         M.Sin
	import . "lib/math"         Sin
*/

//import "fmt" // import packages one by one
//import "os"

import ( // grouped import syntax
	"fmt"
	"os"
)

func main() {
	fmt.Println()
	os.Exit(0)
}
```

**[⬆ Back to top](#contents)**

## Goroutines

> KEYWORDS: **go**.

A "go" statement starts the execution of a function call as an independent
concurrent thread of control, or *goroutine*, within the same address space.

---

[go.go](./src/go.go)

```go
package main

import (
	"fmt"
	"time"
)

// [go]
func printNums(from, to int) {
	for i := from; i <= to; i++ {
		time.Sleep(50 * time.Millisecond)
		fmt.Print(i, " ")
	}
}

func main() {
	fmt.Println("Print nums by Goroutines.")
	go printNums(0, 5)
	go printNums(6, 10)
	// No output, because the main() does not wait, exits before the Goroutines.
	// It is a simple solution to put a wait in the main thread:
	time.Sleep(500 * time.Millisecond)

	// Anonymous Goroutine
	go func() {
		for i := 11; i < 16; i++ {
			fmt.Print(i, " ")
			time.Sleep(100 * time.Millisecond)
		}
	}()
	time.Sleep(1 * time.Second)
}
```

**[⬆ Back to top](#contents)**

### Channels

> KEYWORDS: **chan**.

A channel provides a *mechanism* for concurrently executing functions to
communicate by *sending* and *receiving* values of a specified element type.

A new, initialized channel value can be made using the built-in function `make`,
which takes the channel *type* and an *optional capacity* as arguments (size of
the *buffer*).

A channel may be closed with the built-in function `close`.

#### Deadlock

A deadlock happens when a group of goroutines are waiting for each other and
none of them is able to proceed. The program will get stuck on the channel send
operation waiting forever for someone to read the value.

Deadlock occurs:

1. Receiving Channels More Than Expected
2. Sending Channels More Than Expected

The solution to resolve Channel Deadlock:  
`Number Of Sending Channels = Number Of Receiving Channels`  

If channel is an unbuffered channel, then sending will be blocked until it is
ready to receive. It is easier if we use a buffered channel, which
accepts a limited number of values without any blocking.

---

[chan.go](./src/chan.go)

```go
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
```

**[⬆ Back to top](#contents)**

### Select of Communications

> KEYWORDS: **select**.

A "select" statement chooses which of a set of possible *send* or *receive*
operations will proceed. It looks similar to a "switch" statement but with the
cases all referring to communication operations.

---

[select.go](./src/select.go)

```go
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
```

**[⬆ Back to top](#contents)**

## Used and proposed sources

- [How to Write Go Code](https://go.dev/doc/code)
- [A Tour of Go](https://go.dev/tour/)
- [Effective Go](https://go.dev/doc/effective_go)
- [The Go Programming Language Specification](https://go.dev/ref/spec)
- [Go Bootcamp](https://www.golangbootcamp.com/book)
- [How To Code in Go](https://www.digitalocean.com/community/books/how-to-code-in-go-ebook)
- [An Introduction to Programming in Go](https://www.golang-book.com/books/intro)
- [Practical Go Lessons](https://www.practical-go-lessons.com/)
- [All About GoLang](https://golangdocs.com/)
- [Golang Advanced Tutorial](https://golangbyexample.com/golang-comprehensive-tutorial/)
- [Programming.Guide/Go](https://programming.guide/go/)
- [Go by Example](https://gobyexample.com/)
- [Go Language Patterns](http://www.golangpatterns.info/)
- [Go on Exercism](https://exercism.org/tracks/go/concepts)
- [Oldbloggers Golang posts](https://www.oldbloggers.com/category/golang/)

**[⬆ Back to top](#contents)**

## Author

**Gábor Imolai**

- [Profile](https://github.com/Imolai "Gábor")
- [Telegram](https://t.me/imolaigabor)

## Support

Contributions, issues, and feature requests are welcome!

Give a ⭐️ if you like this project!

## License

[MIT](LICENSE)

