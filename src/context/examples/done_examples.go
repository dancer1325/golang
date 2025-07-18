package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// Ejemplo 1: WithCancel - canal se cierra cuando se llama cancel()
	fmt.Println("=== Ejemplo 1: WithCancel ===")
	ctx1, cancel1 := context.WithCancel(context.Background())

	go func() {
		<-ctx1.Done()
		fmt.Println("Contexto cancelado:", ctx1.Err())
	}()

	time.Sleep(100 * time.Millisecond)
	cancel1() // Cierra el canal Done()
	time.Sleep(100 * time.Millisecond)

	// Ejemplo 2: WithTimeout - canal se cierra cuando expira el timeout
	fmt.Println("\n=== Ejemplo 2: WithTimeout ===")
	ctx2, cancel2 := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel2()

	select {
	case <-ctx2.Done():
		fmt.Println("Timeout alcanzado:", ctx2.Err())
	case <-time.After(300 * time.Millisecond):
		fmt.Println("No debería llegar aquí")
	}

	// Ejemplo 3: Background - Done() puede retornar nil
	fmt.Println("\n=== Ejemplo 3: Background ===")
	bgCtx := context.Background()
	fmt.Printf("Background Done() channel: %v\n", bgCtx.Done())

	// Ejemplo 4: Múltiples llamadas a Done() retornan el mismo canal
	fmt.Println("\n=== Ejemplo 4: Mismo canal ===")
	ctx4, cancel4 := context.WithCancel(context.Background())
	defer cancel4()

	done1 := ctx4.Done()
	done2 := ctx4.Done()
	fmt.Printf("Mismo canal: %v\n", done1 == done2)
}
