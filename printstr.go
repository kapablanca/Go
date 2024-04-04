package piscine

import "github.com/01-edu/z01"

func PrintStr(s string) {
	aString := []byte(s)

	for _, char := range aString {
		z01.PrintRune(rune(char))
	}
}
