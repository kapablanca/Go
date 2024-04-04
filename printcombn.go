package piscine

import (
	"github.com/01-edu/z01"
)

// PrintCombN prints all possible combinations of n different digits in ascending order
func PrintCombN(n int) {
	for i := 0; i <= 9; i++ {
		generateCombinations(n, i, 1, i)
	}
	z01.PrintRune('\n') // Print a newline character at the end
}

// Recursive function to generate combinations
func generateCombinations(n, start, index, digit int) {
	if index == n {
		z01.PrintRune(rune('0' + digit))
		return
	}

	for i := start + 1; i <= 9; i++ {
		generateCombinations(n, i, index+1, digit*10+i)
		if index == 1 {
			z01.PrintRune(',')
			z01.PrintRune(' ')
		}
	}
}
