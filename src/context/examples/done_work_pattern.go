package main

import (
	"context"
	"fmt"
	"time"
)

// Patrón típico: trabajo cancelable usando Done()
func doWork(ctx context.Context) error {
	for i := 0; i < 10; i++ {
		select {
		case <-ctx.Done():
			fmt.Printf("Trabajo cancelado en iteración %d: %v\n", i, ctx.Err())
			return ctx.Err()
		default:
			fmt.Printf("Trabajando... iteración %d\n", i)
			time.Sleep(200 * time.Millisecond)
		}
	}
	fmt.Println("Trabajo completado")
	return nil
}

func main() {
	// Trabajo que se cancela por timeout
	fmt.Println("=== Trabajo con Timeout ===")
	ctx1, cancel1 := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel1()

	doWork(ctx1)

	fmt.Println("\n=== Trabajo Cancelado Manualmente ===")
	ctx2, cancel2 := context.WithCancel(context.Background())

	go func() {
		time.Sleep(800 * time.Millisecond)
		fmt.Println("Cancelando trabajo...")
		cancel2()
	}()

	doWork(ctx2)
}
