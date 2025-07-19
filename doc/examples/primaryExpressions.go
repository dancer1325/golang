package main

import (
	"fmt"
	"strconv"
)

// Persona es un tipo para demostrar expresiones primarias
type Persona struct {
	Nombre string
	Edad   int
}

// Saludar es un método de Persona
func (p Persona) Saludar() string {
	return "Hola, soy " + p.Nombre
}

// Tipo para demostrar type assertion
type Hablante interface {
	Hablar() string
}

func (p Persona) Hablar() string {
	return p.Saludar() + " y tengo " + strconv.Itoa(p.Edad) + " años"
}

func main() {
	// 1. Operand - identificadores, literales, etc.
	x := 10     // x es un operando
	s := "hola" // s es un operando

	// 2. Selector (PrimaryExpr Selector) - acceso a campos o métodos
	p := Persona{"Ana", 30}
	fmt.Println(p.Nombre)    // p.Nombre es una expresión primaria con selector
	fmt.Println(p.Saludar()) // p.Saludar() es una expresión primaria con selector y argumentos

	// 3. Index (PrimaryExpr Index) - acceso a elementos de array o slice
	numeros := []int{10, 20, 30, 40}
	fmt.Println(numeros[2]) // numeros[2] es una expresión primaria con índice

	// 4. Slice (PrimaryExpr Slice) - obtener subconjunto de slice
	fmt.Println(numeros[1:3])   // numeros[1:3] es una expresión primaria con slice
	fmt.Println(numeros[:2])    // numeros[:2] es una expresión primaria con slice
	fmt.Println(numeros[2:])    // numeros[2:] es una expresión primaria con slice
	fmt.Println(numeros[1:3:4]) // numeros[1:3:4] es una expresión primaria con slice de 3 índices

	// 5. TypeAssertion (PrimaryExpr TypeAssertion) - aserción de tipo
	var h Hablante = p
	persona, ok := h.(Persona) // h.(Persona) es una expresión primaria con aserción de tipo
	fmt.Println(persona.Nombre, ok)

	// 6. Arguments (PrimaryExpr Arguments) - llamada a función
	fmt.Println(suma(5, 3)) // suma(5, 3) es una expresión primaria con argumentos

	// 7. MethodExpr - expresión de método
	saludar := Persona.Saludar // Persona.Saludar es una expresión de método
	fmt.Println(saludar(p))    // Llamada a través de la expresión de método

	// 8. Conversion - conversión de tipo
	var i int = 42
	var f float64 = float64(i) // float64(i) es una expresión de conversión
	fmt.Println(f)
}

func suma(a, b int) int {
	return a + b
}
