package main

import "fmt"

// ❌ NO PERMITIDO: Interface suelta con union types a nivel de paquete
// interface {
//     int | float64 | string
// }

// ✅ PERMITIDO: Como type constraint en generics
func Process[T interface{ int | float64 | string }](value T) {
	fmt.Printf("Processing: %v\n", value)
}

// ✅ PERMITIDO: Como tipo nombrado
type Numeric interface {
	int | float64 | string
}

func main() {
	fmt.Println("=== ¿Por qué no se permite interface suelta con union types? ===")

	// ✅ PERMITIDO: Interface literal en variable local
	var numeric interface {
		int | float64 | string
	}

	numeric = 42
	fmt.Printf("numeric = %v\n", numeric)

	// ✅ PERMITIDO: Interface literal en parámetro de función
	processValue := func(v interface{ int | string }) {
		fmt.Printf("Value: %v\n", v)
	}

	processValue(100)
	processValue("hello")

	// ✅ PERMITIDO: Como type constraint
	Process[int](42)
	Process[string]("hello")

	fmt.Println("\n=== Razones por las que no se permite ===")
	fmt.Println("1. Union types están diseñados para CONSTRAINTS, no para valores")
	fmt.Println("2. No puedes llamar métodos en union types")
	fmt.Println("3. Solo son útiles en contexto de generics")
	fmt.Println("4. Evita confusión entre interfaces tradicionales y constraints")

	// Demostrar limitaciones de union types
	fmt.Println("\n=== Limitaciones de union types ===")

	var unionVar interface{ int | string }
	unionVar = 42

	// ❌ No puedes hacer esto con union types:
	// unionVar.SomeMethod()  // No hay métodos comunes

	// Solo puedes hacer type assertions
	if val, ok := unionVar.(int); ok {
		fmt.Printf("Es int: %d\n", val)
	}
}
