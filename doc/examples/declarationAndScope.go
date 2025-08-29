package main

import (
	"fmt"
)

// 1.1  Constant declarations
const Pi = 3.14159
const (
	StatusOK    = 200
	StatusError = 500
)

// 1.2. Type declarations
// type MyInt int						// 2.2 NOT VALID to declare twice | file & package block					uncomment to see the error
type PersonAnother struct {
	Name string
	Age  int
}

// 1.3 Type parameter declarations (generics)
type StackTwo[T any] struct {
	items []T
}

func (s *StackTwo[T]) Push(item T) {
	s.items = append(s.items, item)
}

// 1.4 Variable declarations
var (
	nameAnother = "Go"
)

// 1.5. Function declarations
func add(a, b int) int {
	return a + b
}

func main() {

	// 1.6 Label declarations (used with goto, break, continue)
Loop:
	for i := 0; i < 3; i++ {
		if i == 1 {
			continue Loop
		}
		fmt.Println(i)
	}

	// 2.1
	//variableNotDeclaredAndAssignedDirectly = 23		// uncomment to see the error

	// 2.3 INVALID to declare twice | SAME block
	var nameDeclaredTwiceInSameBlock = "Alfred"
	//var nameDeclaredTwiceInSameBlock = "AlfredTwo"		// uncomment to see the error
}
