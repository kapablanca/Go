package piscine

import "github.com/01-edu/z01"

/* Function that prints an int passed in parameter.
All possible values of type int have to go through.
Cannot convert to int64 */
func PrintNbr(n int) {
	// Check if number is negative to print the minus sign
	if n < 0 {
		z01.PrintRune('-')
	}

	// Store digits of number in a string variable
	number := ""
	lastDigit := 0

	// Checking if n is 0
	if n == 0 {
		number += string('0')
	}
	// Extracting the last digit and appending it to the number until n = 0
	for n != 0 {

		lastDigit = n % 10
		// Making negative digit into positive
		if lastDigit < 0 {
			lastDigit = -lastDigit
		}

		number += string(rune(lastDigit) + '0')
		// Dealing with the minus sign
		n /= 10
	}

	// Printing the n as string
	for i := len(number) - 1; i >= 0; i-- {
		z01.PrintRune(rune(number[i]))
	}
}
