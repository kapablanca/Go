package main

import (
	"os"
	"piscine"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args[1:]

	// checking first argument to implement capital letters if needeed
	capital := false

	if arguments[0] == "--upper" {
		capital = true
		arguments = arguments[1:]
	}
	// checking each argument that is valid and printing
	for _, nbr := range arguments {
		if !piscine.IsNumeric(nbr) {
			z01.PrintRune(' ')
		} else {
			digit := piscine.Atoi(nbr)

			if digit < 1 || digit > 26 {
				z01.PrintRune(' ')
			} else {
				if capital {
					z01.PrintRune(rune(digit - 1 + 'A'))
				} else {
					z01.PrintRune(rune(digit - 1 + 'a'))
				}
			}
		}
	}
}
