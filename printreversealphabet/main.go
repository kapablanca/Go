package main

import "github.com/01-edu/z01"

// Print the Latil lowercase alphabet in reverse order
func main() {
	alphabetLowercase := "abcdefghijklmnopqrstuvwxyz"

	for i := len(alphabetLowercase) - 1; i >= 0; i-- {
		z01.PrintRune(rune(alphabetLowercase[i]))
	}
	z01.PrintRune('\n')
}
