package piscine

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	if n < 0 {

		z01.PrintRune('-')
		n = -n
	}

	count := 0
	num := n
	div := 1

	for num/10 != 0 {
		num = num / 10
		count++
	}

	for count > 0 {

		div = div * 10
		count--
	}

	dig := 0

	for div > 1 {

		dig = n / div
		n = n % div
		div = div / 10
		z01.PrintRune(rune(dig + 48))

	}

	z01.PrintRune(rune(n + 48))
}
