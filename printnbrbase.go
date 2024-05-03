package piscine

import "github.com/01-edu/z01"

// Checking if base contains + or -
func hasSign(base string) bool {
	for _, char := range base {
		if char == '+' || char == '-' {
			return true
		}
	}
	return false
}

// Checking for uniqueness of characters
func hasUnique(base string) bool {
	for i := 0; i < len(base)-1; i++ {
		for j := i + 1; j < len(base); j++ {
			if base[i] == base[j] {
				return false
			}
		}
	}
	return true
}

// Checking validity of a base
func isValid(base string) bool {
	if len(base) < 2 || hasSign(base) || !hasUnique(base) {
		return false
	}
	return true
}

// Print a slice of runes
func printSlice(slice []rune) {
	for _, char := range slice {
		z01.PrintRune(char)
	}
}

// Function that prints an int in anothe base passed as string parameter.
// The function has to manage negative numbers.
// If the base is not valid, the functions prints NV(Not Valid)
// Validity rules for a base:
// Must contain at least 2 characters
// Each character must be unique
// A base should not contain + or - characters
func PrintNbrBase(nbr int, base string) {
	if !isValid(base) {
		z01.PrintRune('N')
		z01.PrintRune('V')
		return
	}

	n := len(base)
	digits := []rune{}
	var mod int

	negative := false
	if nbr < 0 {
		negative = true
	}
	// Extracting the digits in the new base in reverse order
	for nbr >= n || nbr <= -n {
		mod = nbr % n
		nbr /= n
		if negative {
			mod = -mod
		}
		digits = append(digits, rune(base[mod]))
	}
	if negative {
		nbr = -nbr
	}
	digits = append(digits, rune(base[nbr]))

	// Fixing the final form of the number and printing it
	newNbr := []rune{}
	if negative {
		newNbr = append(newNbr, '-')
	}

	for i := len(digits) - 1; i >= 0; i-- {
		newNbr = append(newNbr, digits[i])
	}
	printSlice(newNbr)
}
