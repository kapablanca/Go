package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args[1:]
// checking first argument to implement capital letters if needeed
	capital :=  false

	if arguments[0] == "--upper" {

		capital = true
		arguments = arguments[1:]

	}
	// checking each argument that is valid and printing
	for _, nbr := range arguments {
		if nbr != int || nbr < 1 || nbr > 26 {
			z01.PrintRune(' ')	
		}
	}else if !capital{
		z01.PrintRune(rune(nbr + '`'))
	}else {
		z01.PrintRune(rune(nbr + '@'))
	}
	 
	


}
