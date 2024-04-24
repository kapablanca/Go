package piscine

import (
	"github.com/01-edu/z01"
)

// Prints 'T' if a number is negative, else 'F'.
func IsNegative(nb int) {
	if nb < 0 {
		z01.PrintRune('T')
		z01.PrintRune('\n')
	} else {
		z01.PrintRune('F')
		z01.PrintRune('\n')
	}
}
