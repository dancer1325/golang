package main

import (
	"fmt"
	"runtime"
	"time"
)

var sharedCounter int

func main() {
	fmt.Printf("main goroutine %v\n", getGoroutineID())
	// 1. `go functionCall`
	go func() {
		fmt.Println("executed | goroutine")
	}()

	// 2. `go methodCall`
	go printFunction()

	// 3. `go (Expression)`				wrap with () -> NOT valid
	//go (printFunction())			// uncomment to see the error

	// 4. execution | SAME address
	go increaseCounter("Goroutine-1")
	go increaseCounter("Goroutine-2")

	// 5. independent concurrent thread  == NOT block the main goroutine
	fmt.Printf("main goroutine %v\n", getGoroutineID())

	// 6. NOT ALLOWED | specific built-in functions
	slice := []int{1, 2, 3}
	// go len(slice)
	// go cap(slice)
	fmt.Println(slice) // ONLY, avoid warning NOT use it

	// 7. function's value & parameters are evaluated | main goroutine == BEFORE passing -- to -- NEW goroutine
	printNum(2 + 4)

	// 8. Expression return values are discarded
	//routineId := go getGoroutineID()			// uncomment to see the error
	//printNum(routineId)

	// ⚠️wait for ending up the go routines -- OTHERWISE, function program does NOT wait⚠️
	time.Sleep(500 * time.Millisecond)
	fmt.Printf("Final counter: %d\n", sharedCounter)
}

func printFunction() {
	fmt.Printf("printFunction() \n")
}

func increaseCounter(id string) {
	fmt.Printf("%s - increase counter: %d - goroutine id %d\n", id, sharedCounter, getGoroutineID())
	sharedCounter++
}

// TODO: Why NOT ALWAYS return the SAME values❓
func getGoroutineID() int {
	return runtime.NumGoroutine()
}

func printNum(a int) {
	fmt.Printf("printNum a: %d\n", a)
	fmt.Println("printNum | goroutine ", getGoroutineID())
}
