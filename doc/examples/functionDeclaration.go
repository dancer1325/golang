package main

func main() {

}

// 1. function declarations
// 1.1 simple
func Add(a, b int) int { // signature == (a, b int) int		==  	Parameters [ Result ]
	return a + b
}

// 1.2 with type parameters (== generics)
func Map[T, U any](s []T, f func(T) U) []U { // signature == (s []T, f func(T) U) []U		==  	Parameters [ Result ]
	result := make([]U, len(s))
	for i, v := range s {
		result[i] = f(v)
	}
	return result
}

// 1.3 WITHOUT a body
// func Sort(data []int) // signature == (data []int)		==  	Parameters [ Result ]		// ERROR, because it's valid ONLY | special cases
type Sorter interface {
	Len() int // WITHOUT body | interface declaration
}

// 1.4 if signature has Result 		-> MUST end in terminating statements
func Minus(a, b int) int {
	return a - b
}

// INVALID function			-- Reason: it misses terminating statement --			uncomment to see the error
/*func MissTerminateStatement(s string, r rune) int {
	for i, c := range s {
		if c == r {
			return i
		}
	}
	// invalid: missing return statement
}*/
