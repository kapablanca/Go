package main

import (
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
	z01.PrintRune('\n')
}

func main() {
	arguments := os.Args[1:]

	if len(arguments) < 1 {
		buffer := make([]byte, 1)

		for {
			_, err := os.Stdin.Read(buffer)
			if err != nil {
				break
			}
			PrintString(string(buffer))
		}
		return
	}
	for _, arg := range arguments {
		content, err := os.ReadFile(arg)
		if err != nil {
			ErrorPrint(err)
			os.Exit(1)
		}
		content_string := string(content)
		PrintString(content_string)
	}
}
