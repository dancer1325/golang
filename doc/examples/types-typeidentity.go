package main

import "fmt"

type (
	A0 = []string
	A1 = A0
	A2 = struct{ a, b int }
	A3 = int
	A4 = func(A3, float64) *A0
	A5 = func(x int, _ float64) *[]string

	// identifier Type  -- identifier TypeName
	B0 A0

	// identifier Type  -- identifier TypeLit  -- identifier ArrayType  --
	B1 []string

	// identifier [ TypeParameters ] Type  -- identifier [ TypeParameters ] TypeLit  --  identifier [ TypeParameters ] StructType  --
	D0[P1, P2 any] struct {
		x P1
		y P2
	}
	E0 = D0[int, string]
)

func main() {
	// 1. array types
	var a0 A0
	var a1 A1
	var slice []string

	fmt.Printf("A0 type: %T\n", a0)
	fmt.Printf("A1 type: %T\n", a1)
	fmt.Printf("[]string type: %T\n", slice)

	// 2. struct
	var a2 A2
	var structType struct{ a, b int }

	fmt.Printf("A2 type: %T\n", a2)
	fmt.Printf("struct{ a, b int } type: %T\n", structType)

	// 3. basic types
	var a3 A3
	var intBasic int

	fmt.Printf("A3 type: %T\n", a3)
	fmt.Printf("int type: %T\n", intBasic)

	// 4. function types
	var a4 A4
	var a5 A5

	fmt.Printf("A4 type: %T\n", a4)
	fmt.Printf("A5 type: %T\n", a5)

	// 5. defined types
	// 5.1 B0 A0
	var b0 B0
	var b1 B1
	fmt.Printf("B0 type: %T\n", b0)
	fmt.Printf("B1 type: %T\n", b1)
	fmt.Printf("A0 type: %T\n", a0)

	// 5.1.1 DIFFERENT types
	// var x B0 = a0  		// ERROR
	var x B0 = B0(a0) // OK - conversion
	fmt.Printf("A0 instantiation converted to B0: %T\n", x)
	// var y B1 = a0  		// ERROR
	var y B1 = B1(a0) // OK - conversion
	fmt.Printf("A0 instantiation converted to B1: %T\n", y)

	// 5.2 D0[int, string] E0
	var d0 = D0[int, string]{2, "hello"}
	var e0 E0
	fmt.Printf("D0 type: %T\n", d0)
	fmt.Printf("E0 type: %T\n", e0)
}
