package main

import (
	"os"

	"github.com/01-edu/z01"
)

func printString(s string) {
	for _, char := range s {
		z01.PrintRune(char)
	}
}

// Program that prints the arguments received in the command line
// in ASCII order
func main() {
	arguments := os.Args[1:]
	if len(arguments) < 1 {
		return
	}

	for range arguments {
		for i := 0; i < len(arguments)-1; i++ {
			if arguments[i+1] < arguments[i] {
				arguments[i], arguments[i+1] = arguments[i+1], arguments[i]
			}
		}
	}

	for _, arg := range arguments {
		printString(arg)
		z01.PrintRune('\n')
	}
}
