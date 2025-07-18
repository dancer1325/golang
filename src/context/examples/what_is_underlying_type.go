package main

import "fmt"

type Age int     // Age es un nuevo tipo, pero su underlying type es int
type Name string // Name es un nuevo tipo, pero su underlying type es string

func main() {
	fmt.Println("=== ¿Qué es un underlying type? ===")

	var age Age = 25
	var name Name = "Juan"

	// Age y Name son tipos DIFERENTES
	fmt.Printf("Tipo de age: %T\n", age)   // main.Age
	fmt.Printf("Tipo de name: %T\n", name) // main.Name

	// Pero internamente Go los trata como int y string
	fmt.Println("\n=== Operaciones basadas en underlying type ===")

	// Age puede hacer operaciones de int (su underlying type)
	older := age + Age(5)
	fmt.Printf("Age + 5 = %d (funciona porque underlying type es int)\n", older)

	// Name puede hacer operaciones de string (su underlying type)
	fullName := name + Name(" Pérez")
	fmt.Printf("Name concatenation = %s (funciona porque underlying type es string)\n", fullName)

	// Conversiones explícitas (mismo underlying type)
	var normalInt int = int(age)           // Age -> int
	var normalString string = string(name) // Name -> string

	fmt.Printf("\nConversión Age->int: %d\n", normalInt)
	fmt.Printf("Conversión Name->string: %s\n", normalString)

	// ❌ NO puedes mezclar tipos diferentes (aunque tengan mismo underlying type)
	// var mixed = age + 10  // Error: no puedes sumar Age + int directamente

	// ✅ Necesitas conversión explícita
	var mixed = age + Age(10) // OK: Age + Age
	fmt.Printf("Age + Age(10) = %d\n", mixed)
}
