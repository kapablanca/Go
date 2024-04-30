package piscine

import "github.com/01-edu/z01"

// Function that prints the digits of an int passed in parameter
// in ascending order. All possible values of type int have to go through,
// excluding negative numbers. Conversion to int64 is not allowed.
func PrintNbrInOrder(n int) {
	if n < 0 {
		return
	}

	if n == 0 {
		z01.PrintRune('0')
	} else {
		digits := []int{}
		mod := -1
		// putting the digits of n in the slice digits
		for n != 0 {
			mod = n % 10
			n = n / 10
			digits = append(digits, mod)
		}

		// Reorder the elements of digits in ascending order
		for i := 0; i < len(digits)-1; i++ {
			for j := i + 1; j < len(digits); j++ {
				if digits[j] < digits[i] {
					digits[j], digits[i] = digits[i], digits[j]
				}
			}
		}

		// Printing the number
		for _, digit := range digits {
			z01.PrintRune(rune(digit) + '0')
		}
	}
}
