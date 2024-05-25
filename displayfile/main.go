package main

import (
	"fmt"
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

	// Get file info to know the size of file
	fileInfo, err := file.Stat()
	if err != nil {
		return
	}
	// Reading the contents of the file into a byte slice
	fileSize := fileInfo.Size()
	arr := make([]byte, fileSize)
	file.Read(arr)
	// Saving the contents into a variable and print it
	text := string(arr)
	fmt.Print(text)
}
