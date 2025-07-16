package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	deadline := time.Now().Add(2 * time.Second)
	parentContext := context.Background()
	returnedContext, cancel := context.WithDeadline(parentContext, deadline)
	// 1. returned context != parent context -- Reason:🧠DIFFERENT objects🧠
	fmt.Println("returnedContext == parentContext:", returnedContext == parentContext)

	// 2. returned context == copy of the parent context // TODO:

	// 3. returned context's deadline BEFORE or Equal  `deadline`
	parentContextDeadline, hasParentDeadline := parentContext.Deadline()
	fmt.Printf("parentContextDeadline = %v (has deadline: %v)\n", parentContextDeadline, hasParentDeadline)
	returnedContextDeadline, _ := returnedContext.Deadline()
	fmt.Printf("returnedContextDeadline = %v\n", returnedContextDeadline)
	fmt.Println("returnedContext deadline <= specified deadline:", returnedContextDeadline.Before(deadline) || returnedContextDeadline.Equal(deadline))

	// 4. `Context.Done` channel is closed | 1. deadline expires OR 2. returned `CancelFunc` is called OR 3. parent context's Done channel is closed
	// TODO: from here
	defer cancel() // Siempre llamar cancel para liberar recursos

	select {
	case <-time.After(3 * time.Second):
		fmt.Println("Work completed")
	case <-returnedContext.Done():
		fmt.Println("Deadline exceeded:", returnedContext.Err())
	}
}
