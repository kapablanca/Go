package piscine

/*
Function that returns the last rune of a string
*/
func LastRune(s string) rune {
	// Reverses the string
	s = StrRev(s)
	return FirstRune(s)
}
