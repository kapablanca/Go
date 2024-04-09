package piscine

func IsUpper(s string) bool {
	if len(s) != AlphaCount(s) {
		return false
	}
	for _, letter := range s {
		if letter < 'A' || letter > 'Z' {
			return false
		}
	}
	return true
}
