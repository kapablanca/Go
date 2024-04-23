package piscine

import "github.com/01-edu/z01"

/* Function that prints in ascending order and on a single line
all unique combinations of three different digits, so that the first
digit is lower than the second , and the second is lower than the third
Combinations are seperated by a comma and a spcae */
func PrintComb() {
	// Declaring the variables for each digit of the number and the number
	number := 0
	var firstDigit, secondDigit, thirdDigit int

	for number < 1000 {

		firstDigit = number / 100
		secondDigit = (number % 100) / 10
		thirdDigit = number % 10

		// Print the number if firstDigit < secondDigit < thirdDigit
		if (firstDigit < secondDigit) && (secondDigit < thirdDigit) {

			// Printing each digit by turning it into the rune of the number
			z01.PrintRune(rune(firstDigit) + '0')
			z01.PrintRune(rune(secondDigit) + '0')
			z01.PrintRune(rune(thirdDigit) + '0')

			// Printing the seperators until last number
			if number < 789 {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
		}
		number++
	}
}
