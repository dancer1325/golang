package main

import (
	"fmt"
)

// 1. custom types -- based en -- built-in types
type MyInt int
type MyString string
type MyBool bool

// 2. custom types -- based en -- type literals
type MySlice []int
type MyMap map[string]int
type MyStruct struct{ name string }

func main() {
	// 1. built-in types - underlying type == built-in type
	var normalInt int = 42
	var normalString string = "hello"
	var normalBool bool = true

	fmt.Printf("int's underlying type: %T\n", normalInt)
	fmt.Printf("string's underlying type: %T\n", normalString)
	fmt.Printf("bool's underlying type: %T\n", normalBool)

	// 2. custom types / -- based en -- built-in types,   -> underlying type == base type
	var myInt MyInt = 42
	var myString MyString = "hello"
	var myBool MyBool = true

	fmt.Printf("\nMyInt's type: %T\n", myInt)
	fmt.Printf("MyString's type: %T\n", myString)
	fmt.Printf("MyBool's type: %T\n", myBool)

	// -> 👀you can use underlying type's methods👀
	newNumber := myInt + 4
	fmt.Printf("newNumber value: %v with type %T\n", newNumber, newNumber)
	newNumberReverse := 4 + myInt // ALTHOUGH it's reverse order -> final type == custom type
	fmt.Printf("newNumberReverse value: %v with type %T\n", newNumberReverse, newNumberReverse)
	newString := myString + " world"
	fmt.Printf("newString value: %v with type %T\n", newString, newString)
	newStringReverse := " world" + myString // ALTHOUGH it's reverse order -> final type == custom type
	fmt.Printf("newStringReverse value: %v with type %T\n", newStringReverse, newStringReverse)

	// 3. custom types / -- based en -- type literals,   -> underlying type == literal type
	var slice MySlice = []int{1, 2, 3}
	var mapVar MyMap = map[string]int{"a": 1}
	var structVar MyStruct = struct{ name string }{"test"}

	fmt.Printf("\nMySlice's underlying type: %T\n", slice)
	fmt.Printf("MyMap's underlying type: %T\n", mapVar)
	fmt.Printf("MyStruct's underlying type: %T\n", structVar)

	// -> 👀you can use underlying type's methods👀
	slice = append(slice, 4)
	fmt.Printf("slice value: %v with type %T\n", slice, slice)
}
