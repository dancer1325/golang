package main

import "fmt"

func G() {
	fmt.Println("G: initiate")

	defer D() // function G()	 ==  surrounding function

	fmt.Println("G: PREVIOUS to call function / would trigger panic")
	riskyFunction() // this function can cause panic

	fmt.Println("G ends up")
}

func D() {
	fmt.Println("D: initiate")

	if r := recover(); r != nil {
		fmt.Printf("D: panic captured & managed: %v\n", r)
	} else {
		fmt.Println("D: there was NO panic")
	}
}

func riskyFunction() {
	fmt.Println("riskyFunction: execution")

	panic("something happened")

	fmt.Println("riskyFunction: NEVER reached")
}

func main() {
	// 1. manage panic -- via -- `recover`
	fmt.Println("Example: G defer D / calls `recover`")

	G() // invoke G

	fmt.Println("main: keep on AFTER panic managed") // panic is managed -> can continue

	// 2. panic managed + NOT managed
	func() {
		defer func() { // PREVIOUS line `func()`  ==  surrounding function
			fmt.Println("second panic - defer WITHOUT recover") // reached | END of surrounding function == AFTER deferred function WITH recover
		}()

		defer func() { // PREVIOUS line `func()`  ==  surrounding function
			fmt.Println("second panic - defer WITH & WITHOUT recover") // ALTHOUGH it's out of `recover()` -> ALSO executed
			if r := recover(); r != nil {
				fmt.Printf("second panic - defer WITH & WITHIN recover %v\n", r)
			}
		}()

		panic("panic managed + NOT managed")
	}()

	// 3. panic NOT managed
	func() {
		defer func() { // PREVIOUS line `func()`  ==  surrounding function
			fmt.Println("third panic - defer WITHOUT recover") // reached | END of surrounding function
		}()
		panic("panic NOT managed")
	}()

	fmt.Println("NOT reached, since last panic | this goroutine, NOT managed")
}
