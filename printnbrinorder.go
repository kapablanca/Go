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
		for range digits {
			for j := 0; j < len(digits)-1; j++ {
				if digits[j+1] < digits[j] {
					digits[j], digits[j+1] = digits[j+1], digits[j]
				}
			}
		}

		// Printing the number
		for _, digit := range digits {
			z01.PrintRune(rune(digit) + '0')
		}
	}
}
