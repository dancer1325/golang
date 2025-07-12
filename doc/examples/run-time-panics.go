package main

import "fmt"

func getValue() int {
	return 0 // returns zero at runtime
}

func main() {
	// 1. EXECUTION ERRORS

	fmt.Println("=== 1. Array index out of bounds ===")
	func() {
		defer func() { // 💡defer used | goroutine panic💡
			if r := recover(); r != nil {
				fmt.Println("Panic caught:", r)
			}
		}()

		arr := [3]int{1, 2, 3}
		index := 5              // variable index - no compile-time check
		fmt.Println(arr[index]) // trigger AUTOMATICALLY panic -- Reason:🧠index -- out of -- array🧠
	}()

	fmt.Println("\n=== 2. Nil pointer dereference ===")
	func() {
		defer func() { // 💡defer used | goroutine panic💡
			if r := recover(); r != nil {
				fmt.Println("Panic caught:", r)
			}
		}()

		var ptr *int
		fmt.Println(*ptr) // trigger AUTOMATICALLY panic -- Reason:🧠NOT initialized🧠
	}()

	fmt.Println("\n=== 3. Division by zero ===")
	func() {
		defer func() { // 💡defer used | goroutine panic💡
			if r := recover(); r != nil {
				fmt.Println("Panic caught:", r)
			}
		}()

		x := 10
		y := getValue()    // variable value - no compile-time check
		fmt.Println(x / y) // trigger AUTOMATICALLY panic
	}()

	// 2. EXPLICIT DIRECT PANIC CALLS

	fmt.Println("\n=== 4. Direct panic() call ===")
	func() {
		defer func() { // 💡defer used | goroutine panic💡
			if r := recover(); r != nil {
				fmt.Println("Panic caught:", r)
			}
		}()

		panic("Error") // EXPLICIT panic()
	}()

	fmt.Println("\n=== 5. Panic with runtime.Error ===")
	func() {
		defer func() { // 💡defer used | goroutine panic💡
			if r := recover(); r != nil {
				fmt.Println("Panic caught:", r)
			}
		}()

		panic(fmt.Errorf("runtime error: %s", "something failed")) // panic con error
	}()
}
