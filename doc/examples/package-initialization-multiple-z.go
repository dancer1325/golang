package main

import "fmt"

var secondVar = "Second variable"
var _, initOrder2 = fmt.Println("Init: z_second.go")

func init() {
	fmt.Println("package-initialization-multiple-z - init() de z_second.go")
}

func main() {
	fmt.Println("package-initialization-multiple-z - Variables:", firstVar, secondVar)
}
