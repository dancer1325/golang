package main

import "fmt"

// 1. TYPE SET - Conjunto de tipos que una interface define
type Numeric interface {
	int | float64 | string // Type set: {int, float64, string}
}

type Comparable interface {
	~int | ~string // Type set: {int, string y tipos basados en ellos}
}

// 2. INTERFACE ELEMENTS - Métodos y elementos de tipo
type Writer interface {
	Write([]byte) (int, error) // MethodElem
}

type ReadWriter interface {
	Read([]byte) (int, error)  // MethodElem
	Write([]byte) (int, error) // MethodElem
}

type Constraint interface {
	int | string    // TypeElem (union de type terms)
	~float64        // TypeElem (underlying type)
	String() string // MethodElem
}

// 3. VARIABLES DE INTERFACE - Pueden almacenar valores del type set
type MyInt int
type MyString string

func (m MyInt) String() string    { return fmt.Sprintf("MyInt(%d)", m) }
func (m MyString) String() string { return fmt.Sprintf("MyString(%s)", m) }

// Implementación de Writer
type FileWriter struct{}

func (f FileWriter) Write(data []byte) (int, error) {
	fmt.Printf("Writing: %s\n", string(data))
	return len(data), nil
}

func main() {
	fmt.Println("=== 1. TYPE SET ===")

	// Variables de interface pueden almacenar valores del type set
	var num Numeric

	// ✅ int está en el type set de Numeric
	num = int(42)
	fmt.Printf("num = %v (tipo: %T)\n", num, num)

	// ✅ string está en el type set de Numeric
	num = "hello"
	fmt.Printf("num = %v (tipo: %T)\n", num, num)

	// ✅ float64 está en el type set de Numeric
	num = 3.14
	fmt.Printf("num = %v (tipo: %T)\n", num, num)

	fmt.Println("\n=== 2. UNDERLYING TYPE (~) ===")

	var comp Comparable

	// ✅ int está en el type set
	comp = int(100)
	fmt.Printf("comp = %v\n", comp)

	// ✅ MyInt tiene underlying type int (~int)
	comp = MyInt(200)
	fmt.Printf("comp = %v\n", comp)

	fmt.Println("\n=== 3. INTERFACE ELEMENTS ===")

	// Interface con métodos (MethodElem)
	var writer Writer
	writer = FileWriter{}
	writer.Write([]byte("Hello World"))

	fmt.Println("\n=== 4. VARIABLES NO INICIALIZADAS ===")

	var uninitWriter Writer
	var uninitNumeric Numeric

	fmt.Printf("uninitWriter == nil: %v\n", uninitWriter == nil)
	fmt.Printf("uninitNumeric == nil: %v\n", uninitNumeric == nil)

	fmt.Println("\n=== 5. UNION DE TYPE TERMS ===")

	// TypeElem con múltiples TypeTerms separados por |
	type MultiType interface {
		int | string | bool // Union de 3 type terms
	}

	var multi MultiType
	multi = 42
	fmt.Printf("multi (int): %v\n", multi)
	multi = "text"
	fmt.Printf("multi (string): %v\n", multi)
	multi = true
	fmt.Printf("multi (bool): %v\n", multi)
}
