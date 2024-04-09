package piscine

func IsLower(s string) bool {
	if len(s) != AlphaCount(s) {
		return false
	}
	for _, letter := range s {
		if letter < 'a' || letter > 'z' {
			return false
		}
	}
	return true
}
