package piscine

func IsPrintable(s string) bool {
	for _, letter := range s {
		if letter < ' ' || letter == rune(127) {
			return false
		}
	}
	return true
}
