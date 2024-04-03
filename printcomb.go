package piscine

import "github.com/01-edu/z01"

func PrintComb() {
	i := 48
	for i < 58 {

		j := 48
		for j < 58 {

			if j <= i {
				j++
			}

			k := 48
			for k < 58 {

				if k <= j {
					k++
				}

				z01.PrintRune(rune(i))
				z01.PrintRune(rune(j))
				z01.PrintRune(rune(k))
				z01.PrintRune(44 + 32)

			}

		}

		i++
	}
}
