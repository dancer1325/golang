package main

import "fmt"

// Type constraint (siempre es una interface)
type Numeric interface {
	int | float64 | string
}

// Type parameter T tiene como constraint la interface Numeric
func Process[T Numeric](value T) T {
	// El underlying type de T es la interface Numeric
	fmt.Printf("Procesando valor: %v (tipo: %T)\n", value, value)
	return value
}

// Ejemplo más complejo
type Comparable interface {
	int | string
}

func Compare[T Comparable](a, b T) bool {
	// T es un type parameter
	// Su type constraint es Comparable (una interface)
	// El underlying type de T es la interface Comparable
	return a == b
}

// Constraint con métodos
type Stringer interface {
	String() string
}

type Person struct {
	Name string
}

func (p Person) String() string {
	return p.Name
}

func PrintAny[T Stringer](item T) {
	// T debe implementar Stringer interface
	// El underlying type de T es la interface Stringer
	fmt.Println("Item:", item.String())
}

func main() {
	fmt.Println("=== Type Parameters y Type Constraints ===")

	// T se instancia como int, pero su constraint es Numeric interface
	result1 := Process[int](42)
	fmt.Printf("Resultado int: %d\n", result1)

	// T se instancia como string, pero su constraint es Numeric interface
	result2 := Process[string]("hello")
	fmt.Printf("Resultado string: %s\n", result2)

	// Comparaciones
	fmt.Printf("Compare ints: %v\n", Compare(10, 10))
	fmt.Printf("Compare strings: %v\n", Compare("a", "b"))

	// Con métodos
	person := Person{Name: "Juan"}
	PrintAny(person)

	fmt.Println("\n=== Explicación ===")
	fmt.Println("- T es un type parameter")
	fmt.Println("- Numeric/Comparable/Stringer son type constraints (interfaces)")
	fmt.Println("- El underlying type de T es siempre la interface constraint")
	fmt.Println("- Pero T se instancia con tipos concretos (int, string, etc.)")
}
