package main

import (
	"fmt"
	"log"
)

func main() {
	// 1. label | break's target
Outer:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i == 1 && j == 1 {
				break Outer // use the label
			}
			fmt.Printf("i=%d, j=%d\n", i, j)
		}
	}

	// 2. label | continue's target
Loop:
	// fmt.Println("Loop")			NOT valid // TODO: why?
	for i := 0; i < 5; i++ {
		if i == 2 {
			continue Loop // use the label
		}
		fmt.Printf("i=%d\n", i)
	}

	// 3. Label con goto
	goto Error // use the label
	fmt.Println("dead code")

Error:
	log.Panic("error encountered") // labeled statement
}
