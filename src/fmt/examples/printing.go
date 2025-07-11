package main

import (
	"fmt"
	"math"
)

func main() {
	printGeneralVerbs()

	printBooleanVerbs()

	printIntegerVerbs()

}

func printGeneralVerbs() {
	// 1. `%v`
	var num int = 42
	fmt.Printf("%v\n", num)
	//fmt.Printf(num)				// NOT valid | integer

	// 2. `%+v`
	type Person struct {
		Name string
		Age  int
	}
	person := Person{"Bob", 25}
	fmt.Printf("%v\n", person)  // default one / -- NOT valid | integer
	fmt.Printf("%+v\n", person) // | structs, add ALL fields | ALTHOUGH NOT valid

	// 3. `%#v`
	posInf := math.Inf(1)
	fmt.Printf("%#v\n", posInf)
	negInf := math.Inf(-1)
	fmt.Printf("%#v\n", negInf)
	nan := math.NaN()
	fmt.Printf("%#v\n", nan)

	// 4. `%T`
	var text string = "hello"
	var flag bool = true
	var pi float64 = 3.14

	fmt.Printf("Value: %v, Type: %T\n", num, num)
	fmt.Printf("Value: %v, Type: %T\n", text, text)
	fmt.Printf("Value: %v, Type: %T\n", flag, flag)
	fmt.Printf("Value: %v, Type: %T\n", pi, pi)

	// 5. `%%`
	fmt.Printf("Descuento 25%%\n") // prints %
}

func printBooleanVerbs() {
	// 1. `%t`
	a, b := true, false
	fmt.Printf("a AND b: %t\n", a && b)
	fmt.Printf("a OR b: %t\n", a || b)
}

func printIntegerVerbs() {
	var num int = 42
	// 1. `%b`       2. `%d`
	fmt.Printf("Decimal: %d, Binary: %b\n", num, num)
	fmt.Println(num) // default one
	// 3. `%c`
	fmt.Printf("Unicode: %c\n", num)
	// 4. `%o`     `%O`
	fmt.Printf("Base 8: %o, and Base 8 with 0o prefix:%O\n", num, num)
	// 5. `%q`
	fmt.Printf("1!-quoted character literal / safely escaped -- with -- Go syntax: %q\n", num)
	// 6. `%x`     `%X`
	fmt.Printf("Base 16 / lower-case letters | a-f : %x\n Base 16 / upper-case letters | a-f : %X\n", num, num)
	// 7. `%U`
	fmt.Printf("Unicode format: U+1234 %U\n", num)
}
