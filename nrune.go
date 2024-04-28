package piscine

/*Function that returns the n-th rune of a string.
If not possible, it returns 0*/
func NRune(s string, n int) rune {
	for index, char := range s {
		if index == n-1 {
			return char
		}
	}
	return 0
}
