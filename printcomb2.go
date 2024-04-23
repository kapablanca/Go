package piscine

import "github.com/01-edu/z01"

/*Function that prints in ascending order and on a single line
all possible combinations of two different two-digit numbers.
These combinations are separated by a comma and a space. */
func PrintComb2() {

	// Declaring each number's digits
	var digit1, digit2, digit3, digit4 int

	for firstNumber := 0; firstNumber < 100; firstNumber++ {

		for secondNumber := 0; secondNumber < 100; secondNumber++ {

			digit1 = firstNumber / 10
			digit2 = firstNumber % 10

			digit3 = secondNumber / 10
			digit4 = secondNumber % 10

			// Print when firstNumber < secondNumber, putting space between them and seperator after
			if firstNumber < secondNumber {

				z01.PrintRune(rune(digit1) + '0')
				z01.PrintRune(rune(digit2) + '0')
				z01.PrintRune(' ')
				z01.PrintRune(rune(digit3) + '0')
				z01.PrintRune(rune(digit4) + '0')

				// Skip the seperator in the last combination
				if firstNumber < 98 {

					z01.PrintRune(',')
					z01.PrintRune(' ')
				}
			}
		}
	}
	z01.PrintRune('\n')
}
