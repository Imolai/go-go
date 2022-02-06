# Go, Go!

*Short basic introduction to Go v1.17.6 programming language*  

![Go](https://upload.wikimedia.org/wikipedia/commons/thumb/0/05/Go_Logo_Blue.svg/678px-Go_Logo_Blue.svg.png)  
*Go Logo is Copyright 2018 The Go Authors. All rights reserved.*  

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

- [About](#about)
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
  - [Select of Communications](#select-of-communications)
- [Used and proposed sources](#used-and-proposed-sources)
- [License](#license)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

## Comments

Comments serve as program documentation.

---

[comment.go](./src/comment.go)

```go
[[./src/comment.go]]
```

**[⬆ Back to top](#contents)**

## Literals

A literal is a fixed value assigned to variables or constants.

---

[literal.go](./src/literal.go)

```go
[[./src/literal.go]]
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
[[./src/naming.go]]
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
[[./src/const.go]]
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
[[./src/var.go]]
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
[[./src/type.go]]
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
[[./src/struct.go]]
```

**[⬆ Back to top](#contents)**

#### Anonymous Structs

---

[struct.go](./src/struct_anon.go)

```go
[[./src/struct_anon.go]]
```

**[⬆ Back to top](#contents)**

### Maps

> KEYWORDS: **map**.

A map is an unordered group of *elements* of one type, called the element type,
indexed by a set of unique *keys* of another type, called the key type.

---

[map.go](./src/map.go)

```go
[[./src/map.go]]
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
[[./src/func.go]]
```

**[⬆ Back to top](#contents)**

#### Anonymous Functions

---

[func_anon.go](./src/func_anon.go)

```go
[[./src/func_anon.go]]
```

**[⬆ Back to top](#contents)**

#### Goto in a Function

> KEYWORDS: **goto**.

A "goto" statement transfers control to the statement with the corresponding
label within the same function.

---

[goto.go](./src/goto.go)

```go
[[./src/goto.go]]
```

**[⬆ Back to top](#contents)**

#### Return a Function or Method

> KEYWORDS: **return**.

A "return" statement in a function F terminates the execution of F, and
optionally provides one or more result values.

---

[return.go](./src/return.go)

```go
[[./src/return.go]]
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
[[./src/defer.go]]
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
[[./src/interface.go]]
```

**[⬆ Back to top](#contents)**

## If Conditional Execution

> KEYWORDS: **if**.

"If" statements specify the conditional execution of two branches according to
the value of a boolean expression.

---

[if.go](./src/if.go)

```go
[[./src/if.go]]
```

**[⬆ Back to top](#contents)**

### Else of If

> KEYWORDS: **else**.

If the expression evaluates to true, the "if" branch is executed, otherwise, if
present, the "else" branch is executed.

---

[else.go](./src/else.go)

```go
[[./src/else.go]]
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
[[./src/switch.go]]
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
[[./src/type_switch.go]]
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
[[./src/case.go]]
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
[[./src/default.go]]
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
[[./src/fallthrough.go]]
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
[[./src/for.go]]
```

**[⬆ Back to top](#contents)**

### Break a Loop

> KEYWORDS: **break**.

A "break" statement terminates execution of the innermost "for", "switch", or
"select" statement within the same function.

---

[break.go](./src/break.go)

```go
[[./src/break.go]]
```

**[⬆ Back to top](#contents)**

### Continue a Loop

> KEYWORDS: **continue**.

A "continue" statement begins the next iteration of the innermost "for" loop at
its post statement.

---

[continue.go](./src/continue.go)

```go
[[./src/continue.go]]
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
[[./src/range.go]]
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
[[./src/package.go]]
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
[[./src/import.go]]
```

**[⬆ Back to top](#contents)**

## Goroutines

> KEYWORDS: **go**.

A "go" statement starts the execution of a function call as an independent
concurrent thread of control, or *goroutine*, within the same address space.

---

[go.go](./src/go.go)

```go
[[./src/go.go]]
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
[[./src/chan.go]]
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
[[./src/select.go]]
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

## License

[MIT](LICENSE)

