package piscine

import "github.com/01-edu/z01"

func PrintComb() {
	i := 48
	for i < 58 {

		j := 48
		for j <= i {
			j++
		}
		for j < 58 {

			k := 48
			for k <= j {
				k++
			}

			for k < 58 {

				z01.PrintRune(rune(i))
				z01.PrintRune(rune(j))
				z01.PrintRune(rune(k))
				z01.PrintRune(44 + 32)

				k++

			}

			j++

		}

		i++
	}
}
