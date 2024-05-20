package piscine

import "github.com/01-edu/z01"

// Function that receives a string slice and prints
// each element of the slice in a seperate line
func PrintWordsTables(a []string) {
	for _, word := range a {
		for _, char := range word {
			z01.PrintRune(char)
		}
		z01.PrintRune('\n')
	}
}
