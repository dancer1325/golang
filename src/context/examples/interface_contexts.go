package main

import "fmt"

// ❌ NO PERMITIDO: Interface type literal a nivel de paquete
// interface { Write([]byte) (int, error) }

// ✅ PERMITIDO: Interface type con nombre
type Writer interface {
	Write([]byte) (int, error)
}

func main() {
	fmt.Println("=== ¿Por qué se permite dentro de func pero no fuera? ===")

	// ✅ PERMITIDO: Interface type literal en DECLARACIÓN DE VARIABLE
	var writer interface {
		Write([]byte) (int, error)
	}

	// ✅ PERMITIDO: Interface{} vacía
	var value interface{}

	// ✅ PERMITIDO: Union types en variable
	var numeric interface {
		int | float64 | string
	}

	writer = &FileWriter{}
	value = 42
	numeric = "hello"

	fmt.Printf("writer: %T\n", writer)
	fmt.Printf("value: %T\n", value)
	fmt.Printf("numeric: %T\n", numeric)

	fmt.Println("\n=== Razones de la diferencia ===")
	fmt.Println("1. NIVEL DE PAQUETE: Solo se permiten DECLARACIONES DE TIPOS")
	fmt.Println("2. DENTRO DE FUNC: Se permiten DECLARACIONES DE VARIABLES")
	fmt.Println("3. Interface literal es un TIPO, no una declaración de tipo")

	// Otros contextos donde SÍ se permite
	fmt.Println("\n=== Otros contextos permitidos ===")

	// ✅ En parámetros de función
	processWriter(func(w interface{ Write([]byte) (int, error) }) {
		w.Write([]byte("test"))
	})

	// ✅ En return types
	getWriter := func() interface{ Write([]byte) (int, error) } {
		return &FileWriter{}
	}

	w := getWriter()
	w.Write([]byte("from return"))

	// ✅ En type assertions
	if w2, ok := value.(interface{ String() string }); ok {
		fmt.Printf("Has String method: %v\n", w2)
	}
}

func processWriter(fn func(interface{ Write([]byte) (int, error) })) {
	fn(&FileWriter{})
}

type FileWriter struct{}

func (f *FileWriter) Write(data []byte) (int, error) {
	fmt.Printf("Writing: %s\n", string(data))
	return len(data), nil
}
