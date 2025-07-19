package main

import (
	"fmt"
	"unsafe"
)

// 1. type struct
type emptyStruct struct{}

type structWithSeveralFields struct {
	x, y int
	u    float32
	_    float32 // padding
	A    *[]int
	F    func()
}

// NOT ALLOWED to declare | outside the function body
/*struct {
	x, y int
	u float32
	_ float32  // padding
	A *[]int
	F func()
}*/

type T1 string

type T2 int
type T3 int
type T4 int

// P 		type struct with embedded fields
type P struct {
	T3
	*T4
}

func main() {
	// 1. via struct type
	structType()

	// 2. via struct
	viaStruct()

	// 3. struct -- with -- embedded fields
	structWithEmbeddedFields()

	// 4. struct 	INVALID, because embedded field names are NOT unique
	invalidStruct()
}

func viaStruct() {
	emptyStructWithoutType := struct{}{}
	complexStructWithoutType := struct {
		x, y int
		u    float32
		_    float32 // padding
		A    *[]int
		F    func()
	}{
		x: 10,
		y: 20,
		u: 3.14,
		A: &[]int{1, 2, 3},
		F: func() { fmt.Println("Hello") },
	}

	fmt.Printf("Empty struct size: %d bytes\n", unsafe.Sizeof(emptyStructWithoutType))
	fmt.Printf("Complex struct: x=%d, y=%d\n", complexStructWithoutType.x, complexStructWithoutType.y)
}

func structType() {
	// empty struct
	emptyStruct := emptyStruct{}

	// struct / has 6 fields
	complexStruct := structWithSeveralFields{
		x: 10,
		y: 20,
		u: 3.14,
		A: &[]int{1, 2, 3},
		F: func() { fmt.Println("Hello") },
	}

	fmt.Printf("Empty struct size: %d bytes\n", unsafe.Sizeof(emptyStruct))
	fmt.Printf("Complex struct: x=%d, y=%d\n", complexStruct.x, complexStruct.y)
}

func structWithEmbeddedFields() {
	t4Value := T4(42)

	// P instance
	p := P{
		T3: 10,
		T4: &t4Value,
	}

	structWithEmbeddedFields := struct {
		T1       // embedded field / name == T1
		*T2      // embedded field / name == T2
		P        // embedded field / name == P
		x, y int // field names == x & y
	}{
		T1: "string value",
		T2: new(T2),
		P:  p,
		x:  3,
		y:  2,
	}

	// 👀promoted field👀
	fmt.Printf("access appareantly DIRECT -- to -- embedded's inheirt field: %d & %d\n", structWithEmbeddedFields.T3, *structWithEmbeddedFields.T4)
	fmt.Printf("structWithEmbeddedFields size: %d bytes\n", unsafe.Sizeof(structWithEmbeddedFields))
}

func invalidStruct() {
	/*struct {
		T     		// conflicts with embedded field *T & *P.T
		*T    		// conflicts with embedded field T & *P.T
		*P.T  		// conflicts with embedded field T & *T
	}*/
}
