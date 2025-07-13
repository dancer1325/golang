package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
)

func main() {
	// 1. uses | main function
	ctx := context.Background()

	fmt.Printf("0.1. NON-NIL: ctx == nil? %v\n", ctx == nil)
	fmt.Printf("0.2. EMPTY: ctx.Value(\"key\") = %v\n", ctx.Value("key")) // NO values
	fmt.Printf("0.3. NEVER CANCELED:\n")
	fmt.Printf("   ctx.Err() = %v\n", ctx.Err())
	fmt.Printf("0.4. NO deadline:\n")
	deadline, ok := ctx.Deadline()
	fmt.Printf("deadline = %v\n", deadline)
	fmt.Printf("ok = %v\n", ok)
	fmt.Printf("deadline.IsZero() = %v\n", deadline.IsZero())

	processRequest(ctx, "important data")

	// 2. uses | initialize a server & incoming requests
	startServer()
}

func processRequest(ctx context.Context, data string) {
	fmt.Println("Context.Value(): ", ctx.Value("key"))
	fmt.Printf("Processing: %s\n", data)
}

func startServer() {
	// server -- for the -- context's base
	ctx := context.Background()

	server := &http.Server{
		BaseContext: func(net.Listener) context.Context {
			return ctx
		},
	}

	server.ListenAndServe()
}
