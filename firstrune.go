package piscine

/*
Function that returns the first rune of a string
*/
func FirstRune(s string) rune {
	var first rune
	for _, char := range s {
		first = char
		break
	}
	return first
}
