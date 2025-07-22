package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	// 1. func ReadFull(r Reader, buf []byte) (n int, err error) {}
	readFull()

	// TODO:
}

func readFull() {
	// 1. file
	file, _ := os.Open("file.txt")
	buffer := make([]byte, 100)
	n, err := io.ReadFull(file, buffer) // *os.File				implements 	io.Reader
	fmt.Printf("readFull - file - string(buffer[:n]) %v \n with bytes read %v and error %v\n ", string(buffer[:n]), n, err)

	// 2. buffer bytes
	buf := bytes.NewBuffer([]byte("content"))
	bufferFromBytes := make([]byte, 5)
	nBufferBytes, errBufferBytes := io.ReadFull(buf, bufferFromBytes) // *bytes.Buffer		implements 	io.Reader
	fmt.Printf("readFull - bytes - string(buffer[:n]) %v \n with bytes read %v and error %v\n ", string(bufferFromBytes[:nBufferBytes]), nBufferBytes, errBufferBytes)

	// 3. strings reader
	reader := strings.NewReader("hello world")
	bufferString := make([]byte, 4)
	nBufferString, errBufferString := io.ReadFull(reader, bufferString) // *strings.Reader		implements 	io.Reader
	fmt.Printf("readFull - bytes - string(buffer[:n]) %v \n with bytes read %v and error %v\n ", string(bufferString[:nBufferString]), nBufferString, errBufferString)

	// TODO:	 if `r` returns an error & have read >= `len(buf)` bytes -> error is dropped
}
