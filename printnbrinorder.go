package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	if n < 0 {
		return
	}

	var count_digits [10]int
	digit := 0

	if n == 0 {
		count_digits[0] = 1
	} else {
		for n > 0 {
			digit = n % 10
			count_digits[digit] += 1
			n /= 10
		}
	}

	for i := 0; i < len(count_digits); i++ {
		for j := 0; j < count_digits[i]; j++ {
			z01.PrintRune(rune('0' + i))
		}

	}

}
