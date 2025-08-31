package examples

import (
	"fmt"
	"math/rand"
)

func Hello() {
	fmt.Println("hello")
}

// 2. if you do NOT print -> error
func HelloReturned() string {
	return "hello"
}

func Salutations() {
	fmt.Println("hello, and")
	fmt.Println("goodbye")
}

func Perm() {
	for _, value := range rand.Perm(5) {
		fmt.Println(value)
	}
}

func Add(a, b int) int {
	c := a + b
	fmt.Println(c)
	return c
}
