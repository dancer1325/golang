package main

import (
	"fmt"
	"math"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	// 1. Literal - BasicLit
	literalBasic()

	// 2. Literal - CompositeLit
	literalComposite()

	// 3. Literal - FunctionLit
	literalFunction()

	// 4. OperandName - identifier
	variable := 100
	fmt.Println(variable)

	// 5. OperandName - QualifiedIdent
	constant := math.Pi // math.Pi    == 	packageName.identifier
	fmt.Println(constant)

	// 6. OperandName [ TypeArgs ]
	// TODO:

	// 7. expression / wrapped with ()
	result := (10 + 5) * 2
	fmt.Println(result)
}

func literalFunction() {
	functionAnonymous := func(x, y int) int {
		return x + y
	}

	fmt.Println(functionAnonymous(5, 3))
}

func literalComposite() {
	// Array literal
	arrayLit := [3]int{1, 2, 3}

	// Slice literal
	sliceLit := []string{"uno", "dos", "tres"}

	// Map literal
	mapLit := map[string]int{
		"uno": 1,
		"dos": 2,
	}

	// Struct literal
	personLit := Person{
		Name: "Ana",
		Age:  30,
	}

	fmt.Println(arrayLit, sliceLit, mapLit, personLit)
}

func literalBasic() {
	// int_lit
	integer := 42
	bynary := 0b101010
	octal := 0o52
	hexadecimal := 0x2A
	integerWithSeparator := 1_000_000

	// float_lit
	float := 3.14159
	floatExponential := 1.5e3
	floatWithSeparator := 3.141_592_653_589

	// imaginary_lit
	imaginary := 2i
	imaginaryFloat := 3.14i

	// rune_lit
	runeSimple := 'a'
	runeEscaped := '\n'
	runeUnicode := '世'

	// string_lit
	chainSimple := "hello world"
	chainCrude := `line 1
line 2`

	fmt.Println(integer, bynary, octal, hexadecimal, integerWithSeparator)
	fmt.Println(float, floatExponential, floatWithSeparator)
	fmt.Println(imaginary, imaginaryFloat)
	fmt.Println(runeSimple, runeEscaped, runeUnicode)
	fmt.Println(chainSimple, chainCrude)
}
