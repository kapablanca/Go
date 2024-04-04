package piscine

import (
	"github.com/01-edu/z01"
)

func PrintCombN(n int) {
	generateCombinations("", 0, n, true)
	z01.PrintRune('\n') // Correctly print a newline at the end of function execution
}

func generateCombinations(prefix string, start, n int, isLast bool) {
	if len(prefix) == n {
		for _, r := range prefix {
			z01.PrintRune(rune(r))
		}
		if !isLast {
			z01.PrintRune(',')
			z01.PrintRune(' ')
		}
		return
	}

	for i := start; i <= 9; i++ {
		newPrefix := prefix + intToString(i) // Use the custom int to string conversion
		generateCombinations(newPrefix, i+1, n, i == 9 && len(newPrefix) == n-1)
	}
}

// intToString converts an integer to its string representation.
// This version directly constructs the string expected for single digits, leveraging rune conversion.
func intToString(num int) string {
	return string('0' + rune(num))
}
