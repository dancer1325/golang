package main

import "fmt"

func getA() int {
	fmt.Println("getA()")
	return b + c
}

func getB() int {
	fmt.Println("getB()")
	return 10
}

func getC() int {
	fmt.Println("getC()")
	return 20
}

func getD() int {
	fmt.Println("getD() ")
	return a + 5
}

// 1. declaration order: a, b, c, d
var a = getA() // depends -- on -- b & c	-> ⚠️invoke function postponed⚠️
var b = getB() // INDEPENDENT				-> invoke function
var c = getC() // INDEPENDENT 				-> invoke function
// b & c ALREADY known -> invoke getA()
var d = getD() // depends -- on -- a & a ALREADY known -> invoke getD()

// 2. initialization lifecycle
// if you uncomment it -> program does NOT run
/*var e = f + 1 	// 'a' depends -- on -- 'b'
var f = e + 1 	// 'b' depends -- on -- 'a'*/

// 3. variable declaration initialization / | right-hand side, 1! value
func getValues() (int, string) {
	return 42, "hello"
}

// SIMULTANEOUS initialization
var g, h = getValues()

// 4. _
var config, _ = loadConfig() // blank identifier treated -- as -- any OTHER variable

func loadConfig() (int, string) {
	return 42, "hello"
}

// 5. dependency analysis
// 5.1 lexical reference -- to -- variable or function
var globalVar = 100

func calculate(x int) int {
	return x * 2
}

func main() {
	fmt.Printf("a=%d, b=%d, c=%d, d=%d\n", a, b, c, d)
	fmt.Printf("g=%d, h=%v\n", g, h)
	fmt.Println(config)

	// 5.1
	value := globalVar // lexical reference -- to --  variable != copy by reference -- it's a copy by value --
	fmt.Println("PREVIOUS to update globalVar, value ", value)
	globalVar = 50
	fmt.Println("globalVar ", globalVar)
	fmt.Println("AFTER update globalVar, value ", value) // value NOT change because it's a lexical reference != copy by reference

	result := calculate(50) // reference -- to -- function
	fmt.Println("PREVIOUS to update, result ", result)
	result = calculate(20)
	fmt.Println("AFTER update, result ", result)
}
