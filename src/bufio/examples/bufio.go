package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	// 1. NewReader()
	bufioNewReader := bufio.NewReader(os.Stdin)
	// 1.1 *bufio.Reader
	fmt.Printf("bufioNewReader's type %T\n", bufioNewReader)
	// 1.2 default size
	// bufioNewReader.buf			NOT exported
	fmt.Printf("bufioNewReader's size %v\n", bufioNewReader.Size())
}
