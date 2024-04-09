package piscine

func IsLower(s string) bool {
	if len(s) != AlphaCount(s) {
		return false
	}
	if IsUpper(s) {
		return false
	}
	return true
}
