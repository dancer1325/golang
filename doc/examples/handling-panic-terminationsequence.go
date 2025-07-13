package main

import "fmt"

func main() {
	defer fmt.Println("main: defer executed") // func main()  ==  surrounding function,    		fmt.Println()  ==  invoked function

	fmt.Println("main: call level1")
	level1()
	fmt.Println("main: NEVER executed") // Reason: 🧠level1 throws a panic
}

func level1() {
	defer fmt.Println("level1: defer executed") // func level1()  ==  surrounding function,    		fmt.Println()  ==  invoked function

	fmt.Println("level1: call level2")
	level2()
	fmt.Println("level1: NEVER executed") // Reason: 🧠level2 throws a panic🧠
}

func level2() {
	defer fmt.Println("level2: defer executed") // func level2()  ==  surrounding function,    		fmt.Println()  ==  invoked function

	fmt.Println("level2: call level3")
	level3()
	fmt.Println("level2: NEVER executed") // Reason: 🧠level3 throws a panic🧠
}

func level3() {
	defer fmt.Println("level3: defer executed") // func level3()  ==  surrounding function,    		fmt.Println()  ==  invoked function

	fmt.Println("level3: PREVIOUS to panic")
	panic("unreachable")
	fmt.Println("level3: NEVER executed") // Reason: 🧠PREVIOUS line throws EXPLICITLY a panic🧠
}
