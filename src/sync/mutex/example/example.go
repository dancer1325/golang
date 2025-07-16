package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// 1. 	's 0 value  == unlocked mutex
	zeroValueWasUnlocked()

	// 2. 	== mutual exclusion == 1! goroutine can access -- to -- shared resources
	accessSharedResourcesViaMutex()
	accessSharedResourcesViaWithoutMutex()

	// 3. AFTER FIRST use, NOT copy by value mutex
	afterFirstUseNotCopyByValue()

	// 4. n'th call to `Mutex.Unlock` BEFORE m'th call to `Mutex.Lock`, for any `n < m`
	previousUnLockBeforeNextLockWithoutSpecifyingGoroutineOrder()
	previousUnLockBeforeNextLock()

	// 5. locked Mutex is NOT associated -- with a -- particular goroutine
	lockedMutexNotAssociatedWithOneParticularGoroutine()

	// 6. if mutex is NOT locked & you try to unlock it -> run-time error
	unlockAnUnlockedMutex()

}

func zeroValueWasUnlocked() {
	var mu sync.Mutex // zero-value
	fmt.Printf("mu %d\n", mu)
	if mu.TryLock() { // if it was ALREADY locked -> returns false
		fmt.Println("Mutex zero-value was unlocked")
		mu.Unlock()
	} else {
		fmt.Println("✗ Mutex was ALREADY blocked")
	}
}

func accessSharedResourcesViaMutex() {
	var counter int
	var mu sync.Mutex

	for i := 0; i < 1000; i++ {
		go func() {
			mu.Lock()
			counter++ // Secure access
			mu.Unlock()
		}()
	}

	// sleep to wait for completing the goroutines in the main goroutine
	time.Sleep(time.Second)
	fmt.Println("accessSharedResourcesViaMutex - Counter:", counter) // ALWAYS 10
}

func accessSharedResourcesViaWithoutMutex() {
	var counter int

	// ⚠️if you set up lower number -> consistent, because golang is pretty fast⚠️
	for i := 0; i < 1000; i++ {
		go func() {
			counter++ // NO Secure access
		}()
	}

	// sleep to wait for completing the goroutines in the main goroutine
	time.Sleep(time.Second)
	fmt.Println("accessSharedResourcesViaWithoutMutex - Counter:", counter) // RANDOM output
}

func afterFirstUseNotCopyByValue() {
	var mu sync.Mutex

	// 3.1 BEFORE FIRST use -> FINE to copy
	mu2 := mu
	fmt.Printf("BEFORE FIRST use - mu2 %d\n", mu2)

	// 3.2 AFTER FIRST use
	mu.Lock() // FIRST use
	mu3 := mu // copy by value
	mu.Unlock()
	fmt.Printf("AFTER FIRST use - mu %d\n", mu)
	fmt.Printf("AFTER FIRST use - mu2 %d\n", mu2)
	fmt.Printf("AFTER FIRST use - mu3 %d\n", mu3) // DIFFERENT state
}

func previousUnLockBeforeNextLockWithoutSpecifyingGoroutineOrder() {
	var mu sync.Mutex
	var data int

	// Goroutine A
	go func() {
		mu.Lock()
		data = 42
		mu.Unlock() // n=1 OR m=2 --- NOT guaranteed which goroutine before
	}()

	// Goroutine B  -- NOT guaranteed that internal mu.Lock() is executed BEFORE, because goroutine A could have been triggered BEFORE
	go func() {
		mu.Lock()                                                                           // n=1 OR m=2 --- NOT guaranteed which goroutine before
		fmt.Println("previousUnLockBeforeNextLockWithoutSpecifyingGoroutineOrder - ", data) // NOT consistent
		mu.Unlock()
	}()

	// sleep to wait for completing the goroutines
	time.Sleep(time.Second)
}

func previousUnLockBeforeNextLock() {
	var mu sync.Mutex
	var data int
	var wg sync.WaitGroup // set up goroutine order execution

	wg.Add(1)
	// Goroutine A
	go func() {
		defer wg.Done()
		mu.Lock()
		data = 42
		mu.Unlock() // n=1
	}()

	// Goroutine B
	go func() {
		wg.Wait()                                            // wait for goroutine A completing
		mu.Lock()                                            // m=2
		fmt.Println("previousUnLockBeforeNextLock - ", data) // -> ALWAYS read 42
		mu.Unlock()
	}()

	// sleep to wait for completing the goroutines
	time.Sleep(time.Second)
}

func lockedMutexNotAssociatedWithOneParticularGoroutine() {
	var mu sync.Mutex

	// Goroutine 1: locks the mutex
	go func() {
		fmt.Println("lockedMutexNotAssociatedWithOneParticularGoroutine - Goroutine 1: locks")
		mu.Lock()
	}()

	time.Sleep(500 * time.Millisecond) // ⚠️guarantee goroutine 1 locks FIRST⚠️

	// Goroutine 2: unlocks the mutex / blocked by goroutine 1
	go func() {
		fmt.Println("lockedMutexNotAssociatedWithOneParticularGoroutine - Goroutine 2: unlocks")
		mu.Unlock()
	}()

	time.Sleep(500 * time.Millisecond) // ⚠️guarantee goroutine 2 unlocks⚠️

	fmt.Println("lockedMutexNotAssociatedWithOneParticularGoroutine - mu.TryLock() ", mu.TryLock()) // true == was unlocked
}

func unlockAnUnlockedMutex() {
	var mu sync.Mutex // zero-value is unlocked
	//mu.Unlock()			// run-time error / IMPOSSIBLE to catch  -- uncomment to see it
	if mu.TryLock() {
		fmt.Println("unlockAnUnlockedMutex - it's locked -> you can unlock")
		mu.Unlock()
	} else {
		fmt.Println("unlockAnUnlockedMutex - ALREADY unlocked")
	}
}
