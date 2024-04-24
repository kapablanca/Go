package piscine

import "github.com/01-edu/z01"

/*Function that prints one by one the
characters of a string on the screen */
func PrintStr(s string) {
	for _, char := range s {
		z01.PrintRune(char)
	}
}
