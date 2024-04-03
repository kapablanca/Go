package piscine

import "github.com/01-edu/z01"

func PrintComb2() {
	for i := 0; i <= 99; i++ {
		for j := i + 1; j <= 99; j++ {

			dig1 := i/10 + 48
			dig2 := i%10 + 48

			dig3 := j/10 + 48
			dig4 := j%10 + 48

			z01.PrintRune(rune(dig1))
			z01.PrintRune(rune(dig2))

			z01.PrintRune(' ')

			z01.PrintRune(rune(dig3))
			z01.PrintRune(rune(dig4))

			if i == 98 && j == 99 {
				z01.PrintRune('\n')
			} else {

				z01.PrintRune(',')
				z01.PrintRune(' ')

			}
		}
	}
}
