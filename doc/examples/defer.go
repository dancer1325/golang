package main

import "fmt"

// 1. EXPLICIT `return`
func withReturnStatement() {
	defer fmt.Println("defer / executed | return") // func withReturnStatement()  ==  surrounding function,      fmt.Println()  ==  invoke a function
	fmt.Println("BEFORE return")                   // Reason:🧠`defer ...` is executed | surrounding function executes return🧠
	return
	fmt.Println("NOT reachable code")
}

// 2. reach function's body end
func reachesEnd() {
	defer fmt.Println("defer / executed | function's body end")       // func reachesEnd()  ==  surrounding function,    fmt.Println()  ==  invoke a function
	fmt.Println("function's body end & executed BEFORE `defer ...` ") // function's body end
}

// 3. Goroutine panic
func withPanic() {
	defer fmt.Println("defer / executed | panic") // fmt.Println()		== invoke a function
	fmt.Println("BEFORE panic")
	panic("something wrong") // goroutine panic
	fmt.Println("NOT reachable code")
}

func myFunction() {
	fmt.Println("myFunction")
}

func main() {
	fmt.Println("1. Return statement")
	withReturnStatement()

	fmt.Println("\n 2. Reaches function end")
	reachesEnd()

	fmt.Println("\n 3. Panic (defer still runs)")
	func() {
		defer func() { // defer func(){...}()	== invoke a function,   	PREVIOUS line, `func()`  ==  surrounding function
			if r := recover(); r != nil { // TODO: comprehend from here
				fmt.Println("fixed from the panic", r)
			}
		}()
		withPanic()
	}()

	fmt.Println("\n 4. DeferStmt = \"defer\" Expression")
	func() {
		// 4.1 Expression		== function call
		defer fmt.Println("Expression  ==  function call")

		// 4.2 Expression		== method call
		defer myFunction()

		// 4.3 Expression		wrapped with ()
		//defer (fmt.Println("Expression  wrapped")) 		// uncomment to see error

	}()
}
