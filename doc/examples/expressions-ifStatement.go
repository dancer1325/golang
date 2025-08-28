package main

func f() int {
	return 5
}

var y = 10
var z = 3

func withSimpleStmt() int {
	// x := f()					==		SimpleStmt
	// x < y					== 		Expression
	if x := f(); x < y {
		return x
	} else if x > z {
		return z
	} else {
		return y
	}
}

func main() {
	// 1. with `[ SimpleStmt ";" ]`
	result := withSimpleStmt()
	println("1. with `[ SimpleStmt \";\" ]` ", result) // Output: 5
}
