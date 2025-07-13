package main

import "fmt"

// 1. REPORT run-time panics
func reportRuntimePanic() {
	defer func() { // PREVIOUS line, `func reportRuntimePanic()`  ==  surrounding function
		if r := recover(); r != nil { // recover()		support | report & handle run-time panics
			fmt.Printf("Runtime panic reported: %v\n", r)
		}
	}()

	// Simulate runtime panic
	slice := []int{1, 2, 3}
	index := 5
	fmt.Println(slice[index]) // runtime panic: index out of range
}

// 2. HANDLE program-defined error conditions
func validateAge(age int) {
	defer func() { // PREVIOUS line, `func validateAge(age int)`  ==  surrounding function
		if r := recover(); r != nil { // recover()		report run-time panics
			fmt.Printf("Validation error handled: %v\n", r)
		}
	}()

	if age < 0 {
		panic("age cannot be negative") // program-defined error condition
	}
	if age > 150 {
		panic("age too high") // program-defined error condition
	}

	fmt.Printf("Valid age: %d\n", age)
}

// 3. Custom error handling with recover
func safeExecute(fn func()) (err error) {
	defer func() { // PREVIOUS line, `safeExecute(fn func())`  ==  surrounding function
		if r := recover(); r != nil {
			err = fmt.Errorf("execution failed: %v", r)
		}
	}()

	fn()       // execute POTENTIALLY panicking function
	return nil // if panic happens -> NOT reach this line
}

func main() {
	fmt.Println("=== 1. Report runtime panics ===")
	reportRuntimePanic()

	fmt.Println("\n=== 2. Handle program-defined conditions ===")
	validateAge(-5)
	validateAge(200)
	validateAge(25)

	fmt.Println("\n=== 3. Safe execution wrapper ===")
	err := safeExecute(func() {
		panic("something went wrong")
	})
	if err != nil {
		fmt.Printf("Caught error: %v\n", err)
	}
}
