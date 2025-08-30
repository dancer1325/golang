package main

import "fmt"

// return a `func() int`
func makeCounterClosures() func() int {
	count := 0 // variables defined | surrounding function

	// capture 'count'
	return func() int {
		count++      // get access & modify the variable
		return count // survives
	}
}

func main() {
	// 1. 	"func" Signature FunctionBody()					() execute DIRECTLY
	func(a, b int, z float64) bool { return a*b < int(z) }(2, 3, 4)

	// 2. 	"func" Signature FunctionBody()					assign -- to a -- variable
	literalFunctionAssignedToVariable := func(a, b int, z float64) bool { return a*b < int(z) }
	fmt.Print("literalFunctionAssignedToVariable ", literalFunctionAssignedToVariable)

	// 3. closures
	makeCounterClosures()
}
