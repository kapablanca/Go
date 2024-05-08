package main

import (
	"os"

	"github.com/01-edu/z01"
)

// Changes valid string to rune
func miniAtoi(s string) rune {
	integer := 0
	digit := 0

	for _, char := range s {
		digit = int(char - '0')
		integer = 10*integer + digit
	}
	return rune(integer)
}

// Checks validity of argument
func validAlpha(n rune) bool {
	if n < 0 || n > 25 {
		return false
	}
	return true
}

// Program that prints the corresponding letter in the n position
// of the latin alphabet, where n is each argument received.
// For example 1 matches a, 2 matches b, etc. If n does not match
// a valid position of the alphabet or if the argument is not an integer,
// the program should print a space (" ").
// A flag --upper should be implemented. When used, the program prints the result
// in upper case. The flag will always be the first argument.
func main() {
	// Checking arguments length
	arguments := os.Args[1:]
	if len(arguments) < 1 {
		return
	}
	// Checking for upper flag
	isUpper := false
	if arguments[0] == "--upper" {
		isUpper = true
		arguments = arguments[1:]
	}

	var n rune
	for _, arg := range arguments {
		n = miniAtoi(arg) - 1
		// Not valid char
		if !validAlpha(n) {
			z01.PrintRune(' ')
			// Upper flag active
		} else if isUpper {
			z01.PrintRune(n + 'A')
			// No upper flag
		} else {
			z01.PrintRune(n + 'a')
		}
	}
}
