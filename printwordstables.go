package piscine

import "github.com/01-edu/z01"

// Function that receives a string slice and prints
// each element of the slice in a seperate line
func PrintWordsTables(a []string) {
	for i, word := range a {
		for _, char := range word {
			z01.PrintRune(char)
		}
		// Don't print new line for the last word
		if i != len(a)-1 {
			z01.PrintRune('\n')
		}
	}
}
