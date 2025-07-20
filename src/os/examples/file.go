package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// 1. read file's data
	readFileData()
}

func readFileData() {
	file, err := os.Open("file.go")
	if file != nil {
		fmt.Println("Opened file")
	}
	if err != nil {
		log.Fatal(err) // return values / type error
	}

	data := make([]byte, 100)
	count, err := file.Read(data) // count  == number of bytes read
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("read %d bytes: %q\n", count, data[:count]) // read bytes are stored | data
}
