package main

import "fmt"

var firstVar = "First variable"
var _, initOrder1 = fmt.Println("Init: a_first.go")

func init() {
	fmt.Println("package-initialization-multiple-a - init()")
}
