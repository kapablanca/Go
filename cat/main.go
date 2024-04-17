package main

import (
	"io"
	"os"

	"github.com/01-edu/z01"
)

// Function to print any string in stdout
func PrintString(s string) {
	for _, char := range s {
		z01.PrintRune(char)
	}
}

// Function to print error message
func ErrorPrint(err error) {
	message := "ERROR: " + err.Error()
	PrintString(message)
}

func main() {
	arguments := os.Args[1:]

	if len(arguments) < 1 {
		buffer := make([]byte, 1)
		for {
			_, err := os.Stdin.Read(buffer)
			if err != nil {
				if err == io.EOF {
					continue
				}
				ErrorPrint(err)
				return
			}
			PrintString(string(buffer))
		}

	} else {
		for _, arg := range arguments {
			content, err := os.ReadFile(arg)
			if err != nil {
				ErrorPrint(err)
			}
			content_string := string(content)
			PrintString(content_string)
		}
	}
}
