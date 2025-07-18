package main

import "fmt"

func main() {
	slice := []int{1, 2, 3}

	// ✅ VÁLIDO: Funciones built-in que retornan valores útiles
	go func() {
		fmt.Println("len:", len(slice))
	}()

	go func() {
		ch := make(chan int, 1)
		ch <- 42
		close(ch)
	}()

	// ❌ INVÁLIDO: Built-ins que solo retornan valores sin efectos secundarios
	// go len(slice)        // Error: len(slice) no es un statement válido
	// go cap(slice)        // Error: cap(slice) no es un statement válido
	// go new(int)          // Error: new(int) no es un statement válido

	// ✅ VÁLIDO: Built-ins con efectos secundarios
	ch := make(chan int, 1)
	go func() {
		close(ch) // close() tiene efecto secundario
	}()

	// ✅ VÁLIDO: Built-ins en expresiones más complejas
	go func() {
		_ = len(slice) // Asignación es válida
	}()

	fmt.Println("Programa terminado")
}
