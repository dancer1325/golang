package main

import (
	"fmt"
	"sync"
	"time"
)

var name = "Alfred"

func main() {

	multipleReaderAndOneWriter()

	zeroValueWasUnlocked()

	afterFirstUseNotCopyByValue()

	// TODO: "if lock hold by >=1 readers &"

	// TODO: RLock

	// TODO: RUnlock

}

func multipleReaderAndOneWriter() {
	var rwmu sync.RWMutex // zero-value

	// MULTIPLE SIMULTANEOUS readers
	for i := 1; i <= 3; i++ {
		go reader(i, &rwmu)
	}

	// sleep to wait for completing the goroutines in the main goroutine
	time.Sleep(500 * time.Millisecond)

	// 1 writerRLock (block rest)
	go writer(1, &rwmu)

	// sleep to wait for completing the goroutines in the main goroutine
	time.Sleep(2 * time.Second)

	// ANOTHER readerRLock can NOT access to shared resources
	go reader(1, &rwmu)

	// sleep to wait for completing the goroutines in the main goroutine
	time.Sleep(1 * time.Second)
}

func reader(id int, rwmu *sync.RWMutex) {
	rwmu.RLock()         // MULTIPLE readers can read SIMULTANEOUSLY
	defer rwmu.RUnlock() // executed | end of this execution function

	fmt.Printf("Reader %d: name = %v & name's length %v\n", id, name, len(name))
	time.Sleep(1 * time.Second)
}

func writer(id int, rwmu *sync.RWMutex) {
	rwmu.Lock()         // ⚠️ 1! writerRLock can write & NO readerRLock can access to shared resources
	defer rwmu.Unlock() // executed | end of this execution function

	name = name + " Andrea"
	fmt.Printf("Writer %d: modified the name %v\n", id, name)
	time.Sleep(1 * time.Second)
}

func zeroValueWasUnlocked() {
	var rwmu sync.RWMutex // zero-value
	fmt.Printf("rwmu %d\n", rwmu)
	if rwmu.TryLock() { // if it was ALREADY locked -> returns false
		fmt.Println("Mutex zero-value was unlocked")
		rwmu.Unlock()
	} else {
		fmt.Println("✗ Mutex was ALREADY blocked")
	}
}

func afterFirstUseNotCopyByValue() {
	var rwmu sync.RWMutex

	// 3.1 BEFORE FIRST use -> FINE to copy
	rwmu2 := rwmu
	fmt.Printf("BEFORE FIRST use - rwmu2 %d\n", rwmu2)
	fmt.Printf("BEFORE FIRST use - rwmu2 %d\n", rwmu2)

	// 3.2 AFTER FIRST use
	rwmu.Lock()   // FIRST use
	rwmu3 := rwmu // copy by value
	rwmu.Unlock()
	fmt.Printf("AFTER FIRST use - rwmu %d\n", rwmu)
	fmt.Printf("AFTER FIRST use - rwmu2 %d\n", rwmu2)
	fmt.Printf("AFTER FIRST use - rwmu3 %d\n", rwmu3) // DIFFERENT state
}
