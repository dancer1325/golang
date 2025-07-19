package main

import "fmt"

// 1. DEFINE ALLOWED type parameter's ARGUMENTS
type Numeric interface {
	int | float64 | string
}

// generic function
func PrintValue[T Numeric](value T) { // [T Numeric]  == 	type parameters,			(value T)	== 	type arguments
	fmt.Println("Value:", value)
}

// 2. control the supported operations
type Comparable interface {
	comparable // built-in interface
}

// generic function
func AreEqual[T Comparable](a, b T) bool { // [T Comparable]		== type parameters,			(a, b T)		== type arguments
	return a == b
}

func main() {
	// 1. ALLOWED TYPE ARGUMENTS
	PrintValue(42)      // int 				ALLOWED -- by -- Numeric
	PrintValue(3.14)    // float64  		ALLOWED -- by -- Numeric
	PrintValue("hello") // string	  		ALLOWED -- by -- Numeric
	// PrintValue(true)   		// Error: bool NOT ALLOWED -- by -- Numeric

	// 2. OPERATIONS SUPPORTED
	fmt.Println(AreEqual(5, 5))     // int supports `==`
	fmt.Println(AreEqual("a", "b")) // string soports `==`

	// 3. interface literal form / `interface{ … }` is omitted
	MinFull(1, 2)
	MinOmmited(1, 2)
}

// 3. if interface literal's form == `interface{E}` / `E` == embedded type element -> `interface{ … }` may be omitted

// Definimos una restricción que permite operaciones de ordenamiento (<, >, etc.)
type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64 |
		~string
}

// 3.1 FULL syntax
func MinFull[T interface{ Ordered }](a, b T) T {
	if a < b {
		return a
	}
	return b
}

// 3.2 ommitted syntax
func MinOmmited[T Ordered](a, b T) T { // interface{...}		is omitted
	if a < b {
		return a
	}
	return b
}
