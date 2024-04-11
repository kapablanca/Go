package main

import (
	"os"
	"piscine"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args[1:]
	var letter rune
	// checking first argument to implement capital letters if needeed
	capital := false

	if arguments[0] == "--upper" {

		capital = true
		arguments = arguments[1:]

	}
	// checking each argument that is valid and printing
	for _, nbr := range arguments {
		if !piscine.IsNumeric(nbr) || nbr < "1" || nbr > "26" {
			z01.PrintRune(' ')
		} else if !capital {
			letter = int(nbr)
			z01.PrintRune()
		} else {
			z01.PrintRune(rune)
		}
	}

}
