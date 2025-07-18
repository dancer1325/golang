package main

import "fmt"

// 1. DEFINE ALLOWED TYPE ARGUMENTS
type Numeric interface {
	int | float64 | string // Solo estos tipos están permitidos
}

type Addable interface {
	int | float64 // Solo tipos que se pueden sumar
}

// 2. CONTROLS OPERATIONS SUPPORTED
type Comparable interface {
	int | string // Solo tipos que se pueden comparar con == y <
}

type WithMethods interface {
	String() string // Debe tener método String()
}

// Ejemplos de funciones genéricas
func Process[T Numeric](value T) {
	// T puede ser int, float64, o string (ALLOWED TYPES)
	fmt.Printf("Procesando: %v\n", value)
	// Solo operaciones básicas permitidas
}

func Add[T Addable](a, b T) T {
	// T puede ser int o float64 (ALLOWED TYPES)
	return a + b // Operación + está permitida (SUPPORTED OPERATIONS)
}

func Compare[T Comparable](a, b T) bool {
	// T puede ser int o string (ALLOWED TYPES)
	return a < b // Operación < está permitida (SUPPORTED OPERATIONS)
}

func Print[T WithMethods](item T) {
	// T debe tener método String() (ALLOWED TYPES + SUPPORTED OPERATIONS)
	fmt.Println(item.String()) // Método String() está disponible
}

// Tipo que implementa WithMethods
type Person struct {
	Name string
}

func (p Person) String() string {
	return p.Name
}

func main() {
	fmt.Println("=== 1. DEFINE ALLOWED TYPE ARGUMENTS ===")

	// ✅ Tipos permitidos por Numeric
	Process[int](42)
	Process[float64](3.14)
	Process[string]("hello")

	// ❌ Esto NO compilaría:
	// Process[bool](true)  // bool no está en Numeric interface

	fmt.Println("\n=== 2. CONTROLS OPERATIONS SUPPORTED ===")

	// ✅ Addable permite operación +
	result1 := Add[int](10, 20)
	result2 := Add[float64](1.5, 2.5)
	fmt.Printf("Add int: %d\n", result1)
	fmt.Printf("Add float64: %.1f\n", result2)

	// ✅ Comparable permite operación <
	fmt.Printf("Compare ints: %v\n", Compare[int](5, 10))
	fmt.Printf("Compare strings: %v\n", Compare[string]("a", "b"))

	// ✅ WithMethods permite llamar método String()
	person := Person{Name: "Juan"}
	Print(person)

	fmt.Println("\n=== Explicación ===")
	fmt.Println("Type constraint es una interface que:")
	fmt.Println("1. Define qué tipos están PERMITIDOS (int | string)")
	fmt.Println("2. Controla qué OPERACIONES puedes usar (+, <, métodos)")
}
