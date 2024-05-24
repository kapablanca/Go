package main

import (
	"fmt"
	"io"
	"os"
)

// Program that displays on the standard output the content of a file,
//
//	give as an argument
func main() {
	argument := os.Args[1:]

	if len(argument) < 1 {
		fmt.Print("File name missing")
		return
	}
	if len(argument) > 1 {
		fmt.Print("Too many arguments")
		return
	}
	// Opening the file
	file, err := os.Open(argument[0])
	if err != nil {
		return
	}
	// Make sure the file will close eventually
	defer file.Close()
	// Reading the contents of the file
	content, err := io.ReadAll(file)
	if err != nil {
		return
	}

	text := string(content)
	fmt.Print(text)
}
