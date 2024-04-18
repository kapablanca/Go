package piscine

import "github.com/01-edu/z01"

func PrintTwoDigits(nbr int) {
	digit1 := nbr / 10
	digit2 := nbr % 10

	z01.PrintRune(rune(digit1 + '0'))
	z01.PrintRune(rune(digit2 + '0'))
}

func PrintSeperator() {
	z01.PrintRune(',')
	z01.PrintRune(' ')
}

func DescendComb() {
	for i := 99; i > 0; i-- {
		for j := i - 1; j >= 0; j-- {
			PrintTwoDigits(i)
			z01.PrintRune(' ')
			PrintTwoDigits(j)
			if i != 1 {
				PrintSeperator()
			}
		}
	}
}
