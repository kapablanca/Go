package piscine

func NRune(s string, n int) rune {
	translate := []rune(s)
	if n > 0 && n-1 < len(translate) {
		return translate[n-1]
	}
	return 0
}
