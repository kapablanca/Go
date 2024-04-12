package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args[1:]
	// Stating the variables and arryes I will use
	hasVowel := false
	var indexSlice []int
	var indexArgumentSlice []int
	var vowelSlice []rune
	new_arguments := make([]string, len(arguments))

	// If arguments less than 1, print new line
	if len(arguments) < 1 {

		z01.PrintRune('\n')
	} else {
		// Checking for vowels
		for _, argument := range arguments {

			if isVowel(argument) {
				hasVowel = true
			}
		}
		// If no vowels, print only arguments
		if !hasVowel {
			for _, arg := range arguments {

				printArgs(arg)
				z01.PrintRune('\b')
			}
		} else {
			// Getting the vowels and indexes of the vowels and storing them
			for i, arg := range arguments {
				if isVowel(arg) {

					vowels, positions := checkVowels(arg)

					vowelSlice = append(vowelSlice, vowels...)
					indexSlice = append(indexSlice, positions...)

					// Getting the indexes of the args that each vowel is
					for range vowelSlice {

						indexArgumentSlice = append(indexArgumentSlice, i)
					}

				}

			}
			// Reversing the order of vowels in slice
			for i, j := 0, len(vowelSlice)-1; i < j; i, j = i+1, j-1 {

				vowelSlice[i], vowelSlice[j] = vowelSlice[j], vowelSlice[i]
			}

		}
	}

}

// Returns a slice with the vowels of the string
func checkVowels(phrase string) ([]rune, []int) {

	vowels := [10]rune{'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U'}
	var list_vowels []rune
	var indexes_vowels []int

	// Appending the vowels found and the indexes were they are to slices
	for i, char := range phrase {
		for _, letter := range vowels {

			if char == letter {
				list_vowels = append(list_vowels, letter)
				indexes_vowels = append(indexes_vowels, i)
			}
		}
	}
	return list_vowels, indexes_vowels
}

// Checks if there is a vowel
func isVowel(word string) bool {

	vowels := [10]rune{'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U'}

	for _, char := range word {
		for _, letter := range vowels {

			if char == letter {
				return true
			}
		}
	}
	return false
}

func printArgs(s string) {

	for _, char := range s {
		z01.PrintRune(char)
	}
	z01.PrintRune(' ')
}
