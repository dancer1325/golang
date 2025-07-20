package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	// 1. read file's data
	readFileData("")

	// 2. Write(b []byte) (n int, err error)
	writeFile()

	// 3. io.Writer implemented by File
	ioWriterImplementedByFile()

	// 4. Mkdir example
	createDirectory()
}

func readFileData(fileName string) {
	fileNameToRead := fileName
	if fileNameToRead == "" {
		fileNameToRead = "file.go"
	}
	file, err := os.Open(fileNameToRead)
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

func writeFile() {
	// create a NEW file OR modify EXISTING one
	file, err := os.Create("example.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close() // | exit the function, close the file

	// write bytes DIRECTLY | file
	data := []byte("Hello an example of File.Write\n")
	n, err := file.Write(data) // n == NUMBER of bytes written
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Wrote %d bytes | file\n", n)

	// check written file
	readFileData("example.txt")
}

func ioWriterImplementedByFile() {
	logOutput := os.Stdout // *os.File
	output := io.Writer(logOutput)
	fmt.Printf("writeData - %v\n", output)
}

func createDirectory() {
	dirName := "testdir"

	// create a directory -- with -- permissions 0755 (rwxr-xr-x) (== BEFORE umask)
	err := os.Mkdir(dirName, 0755)
	if err != nil {
		fmt.Printf("Error creating the directory: %v\n", err)
	} else {
		fmt.Printf("Directory with name '%s' created\n", dirName)
	}

	// check directory's information 		-- Reason:🧠check real permissions🧠 --
	info, err := os.Stat(dirName)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Directory permissions: %v\n", info.Mode())

}
