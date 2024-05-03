package main

import (
	"os"

	"github.com/01-edu/z01"
)

// Program that prints the name of the program
func main() {
	arg := os.Args[0]
	lastSlash := -1
	hasExtension := false
	lastDot := -1

	// Extracting the program name from the filepath
	for index, char := range arg {
		if char == '\\' {
			lastSlash = index
		}
		if char == '.' {
			lastDot = index
			hasExtension = true
		}
	}

	if hasExtension {
		arg = arg[lastSlash+1 : lastDot]
	} else {
		arg = arg[lastSlash+1:]
	}

	for _, char := range arg {
		z01.PrintRune(char)
	}
}
