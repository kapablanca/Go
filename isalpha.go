package piscine

// Function that returns true if the string passed as the parameter only
// contains alphanumerical characters or is empty, else returns false
func IsAlpha(s string) bool {
	for _, char := range s {
		if !IsUpper(string(char)) && !IsLower(string(char)) && !isNumber(char) {
			return false
		}
	}
	return true
}
