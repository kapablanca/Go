package piscine

func IsLower(s string) bool {
	if len(s) != AlphaCount(s) {
		return false
	}
	return !IsUpper(s)
}
