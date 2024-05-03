package piscine

// Function that takes two arguments:
// s: a numeric string in a given base
// base: a string representing the base
// and returns the integer value of s in the decimal form.
// The function does not have to manage negative numbers.
// String number must contain only elements that are in base.
// Only valid string numbers will be tested.
// If the base is not valid, it returns 0.
// Validity rules for a base:
// Must contain at least 2 characters
// Each character must be unique
// A base should not contain + or - characters
func AtoiBase(s string, base string) int {
	if !isValid(base) {
		return 0
	}

	n := len(base)
	reverseString := []rune{}
	var number int

	for i := len(s) - 1; i >= 0; i-- {
		reverseString = append(reverseString, rune(s[i]))
	}
	// Transforming number to decimal
	for power, char1 := range reverseString {
		for digit, char2 := range base {
			if char1 == char2 {
				number += digit * IterativePower(n, power)
			}
		}
	}
	return number
}
