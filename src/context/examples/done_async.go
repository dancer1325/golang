package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// Ejemplo de cierre asíncrono del canal Done()
	fmt.Println("=== Cierre Asíncrono del Canal Done() ===")

	ctx, cancel := context.WithCancel(context.Background())

	// Goroutine que escucha el canal Done()
	go func() {
		fmt.Println("Esperando cancelación...")
		<-ctx.Done()
		fmt.Println("¡Canal Done() cerrado! Trabajo cancelado")
	}()

	// Simular trabajo
	fmt.Println("Iniciando trabajo...")
	time.Sleep(500 * time.Millisecond)

	fmt.Println("Llamando cancel()...")
	cancel() // El cierre del canal puede ser asíncrono

	fmt.Println("cancel() retornó inmediatamente")
	time.Sleep(100 * time.Millisecond) // Dar tiempo para que se procese

	fmt.Println("Programa terminado")
}
