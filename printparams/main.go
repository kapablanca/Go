package main

import (
	"os"

	"github.com/01-edu/z01"
)

// Function that prints a string
func printString(s string) {
	for _, char := range s {
		z01.PrintRune(char)
	}
}

// Program that prints the arguments received in command line
func main() {
	arguments := os.Args[1:]
	if len(arguments) < 1 {
		return
	}

	for _, arg := range arguments {
		printString(arg)
		z01.PrintRune('\n')
	}
}
