package main

import "fmt"

func anotherFunction(a int, b int) {
	fmt.Println("a + b ", a+b)
}

func main() {
	// call
	anotherFunction(2, 4)

	// call's arguments restrictions
	// 1. evaluated BEFORE calling the function
	anotherFunction(2, 2+1)
	// 2. arguments MUST be assignable -- to the -- function's parameter types
	//anotherFunction(2, "hello")		// "hello" -- NOT assignable to -- int  		uncomment to see the error
}
