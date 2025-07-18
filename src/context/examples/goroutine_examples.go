package main

import (
	"fmt"
	"time"
)

func main() {
	// Ejemplo 1: Goroutine básica
	go func() {
		fmt.Println("Ejecutándose en goroutine")
	}()

	// Ejemplo 2: Goroutine con función nombrada
	go printNumbers()

	// Ejemplo 3: Múltiples goroutines
	for i := 0; i < 3; i++ {
		go func(id int) {
			fmt.Printf("Goroutine %d ejecutándose\n", id)
		}(i)
	}

	// Esperar para que las goroutines terminen
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Programa principal terminado")
}

func printNumbers() {
	for i := 1; i <= 3; i++ {
		fmt.Printf("Número: %d\n", i)
		time.Sleep(10 * time.Millisecond)
	}
}
