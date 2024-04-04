package piscine

import (
    "github.com/01-edu/z01"
)

func PrintCombN(n int) {
    generateCombinations("", 0, n)
}

func generateCombinations(prefix string, start, n int) {
    if len(prefix) == n {
        for _, r := range prefix {
            z01.PrintRune(rune(r))
        }
        // Determine if this is the last combination to decide whether to print a comma and space
        if prefix < "123456789"[9-n:] {
            z01.PrintRune(',')
            z01.PrintRune(' ')
        } else {
            z01.PrintRune('\n') // Print a newline at the end
        }
        return
    }

    for i := start; i <= 9; i++ {
        newPrefix := prefix + intToString(i) // Use the custom int to string conversion
        generateCombinations(newPrefix, i+1, n)
    }
}

// intToString converts an integer to its string representation.
// This version directly constructs the string expected for single digits, leveraging rune conversion.
func intToString(num int) string {
    return string('0' + rune(num))
}
