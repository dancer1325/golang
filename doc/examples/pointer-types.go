package main

import "fmt"

func main() {
	pointerBasicTypes()

	pointerStructs()

	pointerArrays()

	severalPointersSameType()

	pointerNestedTypes()

	createPointerVianew()

	comparePointerTypes()
}

func pointerBasicTypes() {
	// Base types
	var num int = 42
	var text string = "hello"

	// Pointer types
	var ptrInt *int       // pointer -- to -- int
	var ptrString *string // pointer -- to -- string

	fmt.Printf("ptrInt without initialize: %v\n", ptrInt)       // nil
	fmt.Printf("ptrString without initialize: %v\n", ptrString) // nil

	// 👀assign address👀
	ptrInt = &num
	ptrString = &text

	fmt.Printf("ptrInt assigned: %v\n", ptrInt)
	fmt.Printf("ptrInt point to: %v\n", *ptrInt) // 42
	fmt.Printf("ptrString assigned: %v\n", ptrInt)
	fmt.Printf("ptrString point to: %v\n", *ptrString) // hello
}

type Point struct {
	X, Y int
}

func pointerStructs() {
	// Base type
	var p Point = Point{X: 10, Y: 20}

	// Pointer type
	var ptrPoint *Point // point -- to -- Point

	fmt.Printf("ptrPoint without initialize: %v\n", ptrPoint) // nil

	// assign address
	ptrPoint = &p

	fmt.Printf("ptrPoint assigned -- WRONG format --: %v\n", ptrPoint) // ❌NOT printed address correctly❌
	fmt.Printf("ptrPoint assigned -- RIGHT format --: %p\n", ptrPoint) // print address -- via -- adjusting the print verb
	fmt.Printf("ptrPoint point to: %v\n", *ptrPoint)                   // {10 20}
	fmt.Printf("p.X -- through the -- pointer: %d\n", ptrPoint.X)      // 10
	//fmt.Printf("p.X -- through the -- pointer: %d\n", *ptrPoint.X) // ERROR❌ -- Reason:🧠*ptrPoint is evaluated🧠
}

func pointerArrays() {
	// Base type
	var arr [4]int = [4]int{1, 2, 3, 4}

	// Pointer type
	var ptrArray *[4]int // pointer -- to -- [4]int

	fmt.Printf("ptrArray without initialize: %v\n", ptrArray) // nil

	// assign address
	ptrArray = &arr

	fmt.Printf("ptrArray assigned -- WRONG format --: %v\n", ptrArray) // ❌NOT printed address correctly❌
	fmt.Printf("ptrArray assigned -- RIGHT format --: %p\n", ptrArray) // print address -- via -- adjusting the print verb
	fmt.Printf("ptrArray points to: %v\n", *ptrArray)                  // [1 2 3 4]
	fmt.Printf("arr[0] -- through the -- pointer: %d\n", ptrArray[0])  // 1
}

func severalPointersSameType() {
	// SEVERAL variables / SAME base type
	var a, b, c int = 10, 20, 30

	// SEVERAL pointers type / SAME type (*int)
	var ptr1, ptr2, ptr3 *int

	fmt.Printf("Pointers without initialize: %v, %v, %v\n", ptr1, ptr2, ptr3)

	// assign address / DIFFERENT / EACH one
	ptr1 = &a
	ptr2 = &b
	ptr3 = &c

	fmt.Printf("Pointers assigned / FIRST time, point -- to --: %d, %d, %d\n", *ptr1, *ptr2, *ptr3) // 10, 20, 30

	// 👀reassign address👀
	ptr1 = &b
	fmt.Printf("ptr1 reassigned, points to: %d and ptr2 points to: %d\n", *ptr1, *ptr2)
}

func pointerNestedTypes() {
	var num int = 42
	var ptr *int = &num       // pointer -- to -- int / ALREADY initialized -- via -- assigning address
	var ptrToPtr **int = &ptr // ⚠️pointers -- to -- pointer to int⚠️/ ALREADY initialized -- via -- assigning address

	fmt.Printf("num: %d\n", num)                                                                            // 42
	fmt.Printf("*ptr  == pointer points -- to --: %d\n", *ptr)                                              // 42
	fmt.Printf("**ptrToPtr == pointer points -- to -- another pointer / points -- to --: %d\n", **ptrToPtr) // 42

	fmt.Printf("num's type: %T\n", num)           // int
	fmt.Printf("ptr's type: %T\n", ptr)           // *int
	fmt.Printf("ptrToPtr's type: %T\n", ptrToPtr) // **int
}

func createPointerVianew() {
	// new()
	ptrInt := new(int) //  *int		== pointer -- to -- int
	fmt.Printf("ptrInt's type: %T\n", ptrInt)
	ptrString := new(string) // *string		== pointer -- to -- string
	fmt.Printf("ptrString's type: %T\n", ptrString)
	ptrPoint := new(Point) // *Point		== pointer -- to -- main.Point
	fmt.Printf("ptrPoint's type: %T\n", ptrPoint)

	// 💡| create new -> ALREADY initialized -- with -- zero-values💡
	fmt.Printf("ptrInt WITHOUT initialize: %v\n", ptrInt)         //
	fmt.Printf("*ptrInt WITHOUT initialize: %d\n", *ptrInt)       // 0
	fmt.Printf("ptrString WITHOUT initialize: %v\n", ptrString)   //
	fmt.Printf("*ptrString WITHOUT initialize: %q\n", *ptrString) // ""
	fmt.Printf("ptrPoint WITHOUT initialize: %p\n", ptrPoint)     //
	fmt.Printf("*ptrPoint WITHOUT initialize: %v\n", *ptrPoint)   // {0 0}

	// reassign -- via -- pointers
	*ptrInt = 100
	*ptrString = "modified"
	ptrPoint.X = 5 // ⚠️| structs, you need to use pointer, NOT directly *⚠️

	fmt.Printf("*ptrInt reassigned: %d\n", *ptrInt)       // 100
	fmt.Printf("*ptrString reassigned: %q\n", *ptrString) // "modified"
	fmt.Printf("*ptrPoint reassigned: %v\n", *ptrPoint)   // {5 0}
}

func comparePointerTypes() {
	var a, b int = 10, 10
	var ptr1, ptr2, ptr3 *int // pointers -- to -- int

	// assign address
	ptr1 = &a
	ptr2 = &b
	ptr3 = &a // point -- to -- ALSO a

	// compare pointers -- via -- `==`
	fmt.Printf("ptr1 == ptr2: %t\n", ptr1 == ptr2) // false (== DIFFERENT objects)
	fmt.Printf("ptr1 == ptr3: %t\n", ptr1 == ptr3) // true (SAME object)
	fmt.Printf("ptr1 == nil: %t\n", ptr1 == nil)   // false

	var nilPtr *int
	fmt.Printf("nilPtr == nil: %t\n", nilPtr == nil) // true
}
