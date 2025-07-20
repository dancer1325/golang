package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// 1. Go-like | error handling
	file, err := os.Open("fileNoExisting.go")
	if file != nil {
		fmt.Println("Opened file")
	}
	if err != nil {
		log.Fatal(err) // return values / type error
	}
}
