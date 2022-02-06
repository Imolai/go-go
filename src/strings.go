package main

import (
	"fmt"
	"strings"
)

func main() {
  // A string literal is defined between double quotes.
  const name = "Amy"
  fmt.Println(name)
  const men = "Jack\nJoe\nJohn\nJeremy"
  fmt.Println(men)
  // Raw string literals use backticks (`) instead of double quotes and are
  // interpreted literally, there is no need to escape characters or newlines.
  const daltons = `Joe
William
Jack
Averell`
  fmt.Println(daltons)

  food := "taco"
  fmt.Sprintf("Bring me a %s", food)
  number := 4.3242
  fmt.Sprintf("%.2f", number)
  // The Sprintf verbs:
  // %v	the value in a default format
	// when printing structs, the plus flag (%+v) adds field names
  // %#v	a Go-syntax representation of the value
  // %T	a Go-syntax representation of the type of the value
  // %%	a literal percent sign; consumes no value

  // The default format for %v is:
  //  bool:                    %t
  //  int, int8 etc.:          %d
  //  uint, uint8 etc.:        %d, %#x if printed with %#v
  //  float32, complex64, etc: %g
  //  string:                  %s
  //  chan:                    %p
  //  pointer:                 %p

  // See: https://pkg.go.dev/fmt

  // The strings package
	fmt.Println(strings.Contains("Playground", "ground"))
	fmt.Println(strings.Index("chicken", "ken"))
	fmt.Println(strings.HasPrefix("Gopher", "Go"))
	fmt.Println(strings.HasSuffix("Amigo", "go"))
	fmt.Println("ba" + strings.Repeat("na", 2))
	fmt.Println(strings.Count("cheese", "e"))
	fmt.Println(strings.Replace("oink oink oink", "k", "ky", 2))
	fmt.Println(strings.ReplaceAll("oink oink oink", "oink", "moo"))
	fmt.Println(strings.Trim("¡¡¡Hello, Gophers!!!", "!¡"))
	fmt.Println(strings.TrimRight("¡¡¡Hello, Gophers!!!", "!¡"))
	fmt.Println(strings.TrimLeft("¡¡¡Hello, Gophers!!!", "!¡"))
	fmt.Println(strings.TrimSpace(" \t\n Hello, Gophers \n\t\r\n"))
	fmt.Println(strings.ToLower("Gopher"))
	fmt.Println(strings.Title("her royal highness"))
	fmt.Println(strings.ToUpper("Gopher"))
	s := []string{"The", "Go", "Playground"}
	fmt.Println(strings.Join(s, ", "))
	fmt.Printf("%q\n", strings.Split("a,b,c", ","))
	fmt.Printf("Fields are: %q\n", strings.Fields("  foo bar  baz   "))
}

