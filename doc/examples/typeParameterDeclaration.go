package main

import "fmt"

// 1. TypeParameters | generic functions
// 1.1   1 type parameter
func Print[T any](value T) { //	[T any]		== 		Type parameter,		(value T)	==	 function parameter list
	// T		== IdentifierList,		any		== TypeConstraint
	fmt.Printf("1 type parameter %v\n", value)
}

// 1.2   MULTIPLE type parameter
func MapKeys[K comparable, V any](m map[K]V) []K { //	[K comparable, V any]		MULTIPLE type parameter,		(m map[K]V)   ==	 function parameter list
	// K		== IdentifierList,		comparable		== TypeConstraint
	// V		== IdentifierList,		any		== TypeConstraint
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

// 2. TypeParameters | generic types
// 2.1   1 type parameter
type Stack[T any] struct { // [T any]		==		Type parameter
	// T		== IdentifierList,		any		== TypeConstraint
	items []T
}

func (s *Stack[T]) Push(value T) {
	s.items = append(s.items, value)
}

// 2.2   MULTIPLE type parameter
type Cache[K comparable, V any] struct { // [K comparable, V any]			== 		Type parameter
	// K		== IdentifierList,		comparable		== TypeConstraint
	// V		== IdentifierList,		any		== TypeConstraint
	data map[K]V
}

// 3.	`IdentifierList` non-blank name, MUST be UNIQUE			uncomment to see the error
/*type Cache[K comparable, K any] struct { // [K comparable, V any]			== 		Type parameter
	// K		== IdentifierList,		comparable		== TypeConstraint
	// ⚠️INVALID use K MULTIPLE times as type parameter⚠️
	data map[K]K
}*/

func (c *Cache[K, V]) Set(key K, value V) {
	if c.data == nil {
		c.data = make(map[K]V)
	}
	c.data[key] = value
}

func main() {
	// 1. instantiate generic functions
	instantiateGenericFunctions()

	// 2. instantiate generic types
	instantiateGenericType()
}

func instantiateGenericType() {
	// 1. 		1 type parameter
	var stringStack Stack[string] // explicitly	 				 T = string
	stringStack.Push("a")
	stringStack.Push("b")
	stringStack.Push("c")
	fmt.Printf("instantiateGenericType - 1 type parameters %v\n", stringStack)

	// 2. 		MULTIPLE type parameters
	var userActive Cache[int, bool] // explicitly	 				 K = int, V = bool
	userActive.Set(1, true)
	userActive.Set(2, false)
	fmt.Printf("instantiateGenericType - MULTIPLE type parameters %v\n", userActive)
}

func instantiateGenericFunctions() {
	// 1 type parameter
	Print[int](42) // explicitly	 				 T = int
	Print("hello") // implicitly  (== inferred)	 T = string

	// MULTIPLE type parameters
	users := map[string]int{"Alice": 25, "Bob": 30}
	keys := MapKeys[string, int](users) // explicitly 					K = string, V = int
	keys2 := MapKeys(users)             // implicitly  (== inferred) 	K = string, V = int
	fmt.Println(keys, keys2)
}
