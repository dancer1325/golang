package main

import (
	"fmt"
	"time"
)

var sharedCounter int // Mismo espacio de direcciones

func main() {
	fmt.Println("=== Demostración de goroutines ===")

	// 1. Función normal (secuencial)
	fmt.Println("Llamada normal:")
	normalFunction("A")
	normalFunction("B")

	// 2. Con 'go' (concurrente)
	fmt.Println("\nCon 'go' (concurrente):")
	go goroutineFunction("X") // Independiente
	go goroutineFunction("Y") // Independiente

	// 3. Mismo espacio de direcciones
	fmt.Println("\nMismo espacio de direcciones:")
	go incrementCounter("Goroutine-1")
	go incrementCounter("Goroutine-2")

	// Esperar para ver los resultados
	time.Sleep(200 * time.Millisecond)
	fmt.Printf("Contador final: %d\n", sharedCounter)
}

func normalFunction(id string) {
	fmt.Printf("Función normal %s ejecutándose\n", id)
	time.Sleep(50 * time.Millisecond)
}

func goroutineFunction(id string) {
	fmt.Printf("Goroutine %s ejecutándose independientemente\n", id)
	time.Sleep(50 * time.Millisecond)
}

func incrementCounter(id string) {
	for i := 0; i < 3; i++ {
		sharedCounter++ // Acceso a variable compartida
		fmt.Printf("%s incrementó contador a %d\n", id, sharedCounter)
		time.Sleep(10 * time.Millisecond)
	}
}
