package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// 1. Context WITH deadline
	deadline := time.Now().Add(5 * time.Second)
	ctx, cancel := context.WithDeadline(context.Background(), deadline)
	defer cancel()

	// verify deadline
	if d, ok := ctx.Deadline(); ok {
		fmt.Printf("1. Context WITH deadline -- Deadline: %v\n", d)
		fmt.Printf("1. Context WITH deadline -- Time left: %v\n", time.Until(d))
	} else {
		fmt.Println("1. Context WITH deadline -- No deadline set")
	}

	// 2. Context WITHOUT deadline
	context := context.Background()
	if _, ok := context.Deadline(); !ok {
		fmt.Println("2. Context WITHOUT deadline -- No deadline - can run indefinitely")
	}

}
