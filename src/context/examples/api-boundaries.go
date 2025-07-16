package main

import (
	"context"
	"fmt"
	"time"
)

// ACROSS API BOUNDARIES == 1! process & SEVERAL functions/packages

func main() {
	// TODO: comprehend from here
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	ctx = context.WithValue(ctx, "requestID", "req-123")

	// Context cruza límites de API entre funciones
	handleRequest(ctx)
}

func handleRequest(ctx context.Context) {
	requestID := ctx.Value("requestID")
	fmt.Printf("Handler - Request ID: %s\n", requestID)

	// Context pasa a otra función (API boundary)
	validateUser(ctx)
}

func validateUser(ctx context.Context) {
	requestID := ctx.Value("requestID")
	fmt.Printf("Validator - Request ID: %s\n", requestID)

	// Context pasa a base de datos (API boundary)
	queryDatabase(ctx)
}

func queryDatabase(ctx context.Context) {
	select {
	case <-time.After(1 * time.Second):
		fmt.Println("Database - Query completed")
	case <-ctx.Done():
		fmt.Println("Database - Query cancelled:", ctx.Err())
	}
}
