package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// 1. if context's work is done -> channel is closed
	// 1.1 WithCancel
	fmt.Println("1.1 WithCancel")
	ctx1, cancel1 := context.WithCancel(context.Background())

	go func() {
		<-ctx1.Done()
		fmt.Println("context cancelled:", ctx1.Err())
	}()

	time.Sleep(100 * time.Millisecond)
	cancel1() // | WithCancel, if you call cancel() -> canal `ctx1.Done()` -- TODO: --
	time.Sleep(100 * time.Millisecond)

	// 1.2 WithTimeout
	// if timeout expires -> canal closes
	fmt.Println("\n1.2 WithTimeout ")
	ctx2, cancel2 := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel2()

	select {
	case <-ctx2.Done():
		fmt.Println("Timeout reach:", ctx2.Err())
	case <-time.After(300 * time.Millisecond):
		fmt.Println("NO reached")
	}

	// 1.3 Background
	fmt.Println("\n1.3 Background")
	bgCtx := context.Background()
	fmt.Printf("Background Done() channel: %v\n", bgCtx.Done()) // returns `nil`

	// 2. if you call SUCCESSIVELY Done() -> return the SAME results
	fmt.Println("\n2. if you call SUCCESSIVELY Done() -> return the SAME results")
	ctx4, cancel4 := context.WithCancel(context.Background())
	defer cancel4()

	done1 := ctx4.Done()
	done2 := ctx4.Done()
	fmt.Printf("SAME channel: %v\n", done1 == done2)
}
