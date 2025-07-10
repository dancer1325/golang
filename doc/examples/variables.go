package main

import (
	"errors"
)

// 1. variable declaration
var name string   // Reserve storage -- for a -- string
var age int       // Reserve storage -- for an -- int
var numbers []int // Reserve storage -- for a -- slice

// 2. function declaration
func divide(a, b int) (result int, err error) {
	// 2.1 		a, b 	parameters, reserved storage
	if b == 0 {
		err = errors.New("división por cero")
		return
	}
	// 2.2 		result, err 	result, reserved storage
	result = a / b
	return
}
