package main

import (
	"fmt"
	"time"
)

func expensiveCalculation() int {
	fmt.Println("🔄 Calculando en goroutine principal...")
	time.Sleep(100 * time.Millisecond)
	return 42
}

func getCurrentTime() string {
	return fmt.Sprintf("Tiempo: %v", time.Now().Format("15:04:05.000"))
}

func worker(id string, value int, timestamp string) {
	fmt.Printf("Goroutine %s recibió: value=%d, %s\n", id, value, timestamp)
	time.Sleep(200 * time.Millisecond)
	fmt.Printf("Goroutine %s terminó\n", id)
}

func main() {
	fmt.Println("=== Evaluación de parámetros en calling goroutine ===")

	// Los parámetros se evalúan ANTES de crear la goroutine
	fmt.Println("Antes de 'go worker(...)'")

	go worker(
		"A",
		expensiveCalculation(), // Se ejecuta en main goroutine
		getCurrentTime(),       // Se ejecuta en main goroutine
	)

	fmt.Println("Después de 'go worker(...)' - main continuó inmediatamente")

	// Demostrar que la evaluación es inmediata
	counter := 0
	for i := 0; i < 3; i++ {
		counter++
		go func(val int) {
			time.Sleep(50 * time.Millisecond)
			fmt.Printf("Goroutine recibió counter = %d\n", val)
		}(counter) // counter se evalúa AHORA (1, 2, 3)
	}

	// Comparar con captura por referencia (comportamiento diferente)
	fmt.Println("\n=== Comparación: captura por referencia ===")
	for i := 0; i < 3; i++ {
		counter++
		go func() {
			time.Sleep(100 * time.Millisecond)
			fmt.Printf("Goroutine capturó counter = %d\n", counter) // Valor final
		}()
	}

	time.Sleep(300 * time.Millisecond)
	fmt.Println("Programa terminado")
}
