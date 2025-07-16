package main

import "fmt"

// 1. NOT ALLOWED, define interface | root level
// 1.1 TypeElem.
/*interface {
	int | float64 | string
}*/
// 1.2 MethodElem.
/*interface {
	Write([]byte) (int, error)
}*/

func main() {
	// 2. Variable of interface type
	// 2.1 empty
	var value interface{}

	// include ALL types
	value = 42
	fmt.Printf("value = %v (tipo: %T)\n", value, value)

	value = "hello"
	fmt.Printf("value = %v (tipo: %T)\n", value, value)

	value = []int{1, 2, 3}
	fmt.Printf("value = %v (tipo: %T)\n", value, value)

	fmt.Println("\n=== 2. VARIABLE NO INICIALIZADA ===")

	var uninit interface{}
	fmt.Printf("uninit == nil: %v\n", uninit == nil)

	fmt.Println("\n=== 3. INTERFACE CON MÉTODOS ===")

	// Interface con MethodElem
	var writer interface {
		Write([]byte) (int, error)
	}

	// Implementación
	writer = &FileWriter{}
	writer.Write([]byte("Hello"))

	fmt.Println("\n=== 4. INTERFACE CON TYPE ELEMENTS (Go 1.18+) ===")

	// Interface con TypeElem (union types)
	var numeric interface {
		int | float64 | string
	}

	// Solo puede almacenar tipos del type set
	numeric = 42
	fmt.Printf("numeric = %v\n", numeric)

	numeric = 3.14
	fmt.Printf("numeric = %v\n", numeric)

	numeric = "text"
	fmt.Printf("numeric = %v\n", numeric)

	fmt.Println("\n=== 5. INTERFACE MIXTA ===")

	// Interface con MethodElem + TypeElem
	var constraint interface {
		~int | ~string  // TypeElem
		String() string // MethodElem
	}

	constraint = MyInt(100)
	fmt.Printf("constraint = %v\n", constraint)

	fmt.Println("\n=== 6. INTERFACE ANIDADA ===")

	var complex interface {
		interface{ Read([]byte) (int, error) } // Interface embebida
		Write([]byte) (int, error)             // Método adicional
	}

	complex = &ReadWriter{}
	fmt.Printf("complex type: %T\n", complex)
}

// Tipos de apoyo
type FileWriter struct{}

func (f *FileWriter) Write(data []byte) (int, error) {
	fmt.Printf("Writing: %s\n", string(data))
	return len(data), nil
}

type MyInt int

func (m MyInt) String() string { return fmt.Sprintf("MyInt(%d)", m) }

type ReadWriter struct{}

func (rw *ReadWriter) Read(data []byte) (int, error)  { return 0, nil }
func (rw *ReadWriter) Write(data []byte) (int, error) { return len(data), nil }
