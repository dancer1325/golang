package main

import "fmt"

func main() {
	type S struct {
		a int
		b float64
	}
	variableOfTypeS := new(S)
	fmt.Printf("Type of variableOfTypeS: %T\n", variableOfTypeS)  // *main.S
	fmt.Printf("Value of variableOfTypeS: %p\n", variableOfTypeS) // 0xc000014078 (address)
	fmt.Printf("Pointed value: %+v\n", *variableOfTypeS)          // {a:0 b:0}
}
