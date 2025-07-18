package main

import (
	"fmt"
	"time"
)

func calculate(x, y int) int {
	result := x + y
	fmt.Printf("Calculando %d + %d = %d\n", x, y, result)
	return result
}

func multipleReturns() (int, string, bool) {
	fmt.Println("Función con múltiples valores de retorno")
	return 42, "hello", true
}

func main() {
	fmt.Println("=== Valores de retorno descartados en goroutines ===")

	// ✅ Llamada normal - puedes capturar el valor
	result := calculate(5, 3)
	fmt.Printf("Resultado capturado: %d\n", result)

	// ❌ Con 'go' - NO puedes capturar el valor de retorno
	// result2 := go calculate(10, 20)  // Error de compilación

	// ✅ Con 'go' - el valor se descarta automáticamente
	go calculate(10, 20) // El resultado (30) se pierde

	// ✅ Múltiples valores de retorno también se descartan
	go multipleReturns() // Los valores (42, "hello", true) se pierden

	// Si necesitas el resultado, usa canales
	fmt.Println("\n=== Alternativa: usar canales para capturar resultados ===")
	resultChan := make(chan int)

	go func() {
		result := calculate(15, 25)
		resultChan <- result // Enviar resultado por canal
	}()

	// Recibir el resultado
	goroutineResult := <-resultChan
	fmt.Printf("Resultado de goroutine via canal: %d\n", goroutineResult)

	time.Sleep(100 * time.Millisecond)
	fmt.Println("Programa terminado")
}
