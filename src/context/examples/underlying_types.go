package main

import (
	"fmt"
	"reflect"
)

// Tipos personalizados basados en tipos predeclarados
type MyInt int
type MyString string
type MyBool bool

// Tipos personalizados basados en type literals
type MySlice []int
type MyMap map[string]int
type MyStruct struct{ name string }

func main() {
	fmt.Println("=== Underlying Types ===")

	// 1. Tipos predeclarados - underlying type es el mismo tipo
	var normalInt int = 42
	var normalString string = "hello"
	var normalBool bool = true

	fmt.Printf("int underlying type: %v\n", reflect.TypeOf(normalInt))
	fmt.Printf("string underlying type: %v\n", reflect.TypeOf(normalString))
	fmt.Printf("bool underlying type: %v\n", reflect.TypeOf(normalBool))

	// 2. Tipos personalizados - underlying type es el tipo base
	var myInt MyInt = 42
	var myString MyString = "hello"
	var myBool MyBool = true

	fmt.Printf("\nMyInt type: %v\n", reflect.TypeOf(myInt))
	fmt.Printf("MyString type: %v\n", reflect.TypeOf(myString))
	fmt.Printf("MyBool type: %v\n", reflect.TypeOf(myBool))

	// 3. Type literals - underlying type es el literal mismo
	var slice []int = []int{1, 2, 3}
	var mapVar map[string]int = map[string]int{"a": 1}
	var structVar struct{ name string } = struct{ name string }{"test"}

	fmt.Printf("\n[]int underlying type: %v\n", reflect.TypeOf(slice))
	fmt.Printf("map[string]int underlying type: %v\n", reflect.TypeOf(mapVar))
	fmt.Printf("struct underlying type: %v\n", reflect.TypeOf(structVar))

	// 4. Demostrar que comparten el underlying type
	fmt.Println("\n=== Conversiones (mismo underlying type) ===")

	// Conversión entre tipos con mismo underlying type
	normalInt = int(myInt)   // MyInt -> int
	myInt = MyInt(normalInt) // int -> MyInt

	fmt.Printf("Conversión exitosa: %d <-> %d\n", normalInt, myInt)

	// 5. Operaciones permitidas por underlying type
	fmt.Println("\n=== Operaciones basadas en underlying type ===")

	// MyInt puede usar operaciones de int
	result := myInt + MyInt(10)
	fmt.Printf("MyInt arithmetic: %d + 10 = %d\n", myInt, result)

	// MyString puede usar operaciones de string
	myStr := MyString("Hello") + MyString(" World")
	fmt.Printf("MyString concatenation: %s\n", myStr)
}
