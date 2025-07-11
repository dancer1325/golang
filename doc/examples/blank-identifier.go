package main

import (
	"errors"
	"fmt"
)

func divideTwo(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("división por cero")
	}
	return a / b, nil
}

func main() {
	// 1. anonymous placeholder
	var resultTwo, _ = divideTwo(10, 2)
	fmt.Println("divideTwo(10,2) ", resultTwo)
	// 2. assignment statements
	var resultThree int
	resultThree, _ = divideTwo(10, 1)
	fmt.Println("divideTwo(10,1) ", resultThree)
}
