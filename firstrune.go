package piscine

/*
Function that returns the first rune of a string
*/
func FirstRune(s string) rune {
	for _, char := range s {
		return char
	}
}
