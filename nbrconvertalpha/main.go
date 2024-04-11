package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args[1:]

	capital := false
	if arguments[0] == "--upper" {
		capital = true
		arguments = arguments[1:] // Adjust arguments to skip the --upper flag
	}

	for _, arg := range arguments {
		if !IsNumeric(arg) {
			z01.PrintRune(' ')
		} else {
			digit := Atoi(arg)
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
	z01.PrintRune('\n')
}

// IsNumeric checks if a string is a valid positive integer.
func IsNumeric(s string) bool {
	for _, c := range s {
		if c < '0' || c > '9' {
			return false
		}
	}
	return true
}

// Atoi converts a string to an integer.
func Atoi(s string) int {
	n := 0
	for _, c := range s {
		n = n*10 + int(c-'0')
	}
	return n
}
