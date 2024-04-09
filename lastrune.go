package piscine

func LastRune(s string) rune {
	translate := []rune(s)
	return translate[len(translate)-1]
}
