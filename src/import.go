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
