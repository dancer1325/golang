package main

import "fmt"

// declared | package block
func init() {
	fmt.Println("FIRST init function")
}

// MULTIPLE init function are ALLOWED
func init() {
	fmt.Println("SECOND init function")
}

// package variables initialized BEFORE init functions
var initializationVariable = getGlobalVar()

func getGlobalVar() int {
	fmt.Println("getGlobalVar()")
	return 10
}

func main() {
	// NOT ALLOWED to refer
	// init()
	// var f = init

	fmt.Println("main()")
	fmt.Println(initializationVariable)
}
