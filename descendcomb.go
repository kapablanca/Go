package piscine

import "github.com/01-edu/z01"

// Function that prints in descending order and on a single line all
//  possible combinations of two different two-digit numbers. These
//  combinations are separated by a commma and a space
func DescendComb() {
	var digit1, digit2, digit3, digit4 rune

	for i := 99; i > 0; i-- {
		for j := i - 1; j >= 0; j-- {
			// First number digits
			digit1 = rune(i/10) + '0'
			digit2 = rune(i%10) + '0'
			// Second number digits
			digit3 = rune(j/10) + '0'
			digit4 = rune(j%10) + '0'
			// Printing numbers
			z01.PrintRune(digit1)
			z01.PrintRune(digit2)
			z01.PrintRune(' ')
			z01.PrintRune(digit3)
			z01.PrintRune(digit4)
			// Printing seperator
			if i != 0 {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
		}
	}
}
