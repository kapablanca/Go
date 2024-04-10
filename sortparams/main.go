package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args[1:]

	for i := 0; i < len(arguments); i++ {
		for j := 0; j < len(arguments)-1; j++ {
			if arguments[j] > arguments[j+1] {
				arguments[j], arguments[j+1] = arguments[j+1], arguments[j]
			}
		}
	}

	for _, word := range arguments {
		for _, character := range word {
			z01.PrintRune(character)
		}
		z01.PrintRune('\n')
	}
}
