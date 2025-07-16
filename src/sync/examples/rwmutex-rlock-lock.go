package main

import (
	"fmt"
	"sync"
	"time"
)

var counter = 0
var rwmu sync.RWMutex

func main() {
	fmt.Println("=== WITHOUT RLock - RACE CONDITION ===")
	go unsafeReader()
	go slowWriter()
	time.Sleep(3 * time.Second)

	counter = 0 // reset

	fmt.Println("\n=== WITH RLock - THREAD SAFE ===")
	go safeReader()
	go slowWriterSafe()
	time.Sleep(3 * time.Second)
}

// modify shared resources slowly
func slowWriter() {
	for i := 1; i <= 10; i++ {
		counter = i * 100 // 100, 200, 300...
		time.Sleep(200 * time.Millisecond)
	}
}

func slowWriterSafe() {
	for i := 1; i <= 10; i++ {
		rwmu.Lock()
		counter = i * 100 // CONSISTENT values / ONLY 1! can write | SAME time
		rwmu.Unlock()
		time.Sleep(200 * time.Millisecond)
	}
}

func unsafeReader() {
	for i := 0; i < 20; i++ {
		fmt.Printf("Unsafe: counter = %d\n", counter) // ⚠️INCONSISTENT values / EACH program execution⚠️
		time.Sleep(100 * time.Millisecond)
	}
}

func safeReader() {
	for i := 0; i < 20; i++ {
		rwmu.RLock()
		fmt.Printf("Safe: counter = %d\n", counter) // MANY ALLOWED readers & CONSISTENT values
		rwmu.RUnlock()
		time.Sleep(100 * time.Millisecond)
	}
}
