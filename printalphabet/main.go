package main

import "github.com/01-edu/z01"

// Function to print the Latin alphabet in lowercase letters
func main() {
	alphabetLowercase := "abcdefghijklmnopqrstuvwxyz"

	for _, letter := range alphabetLowercase {
		z01.PrintRune(letter)
	}
	z01.PrintRune('\n')
}
