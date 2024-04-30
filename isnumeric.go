package piscine

// Function that returns true if the string passed as parameter
// contains only numerical characters, else returns false
func IsNumeric(s string) bool {
	for _, char := range s {
		if char < '0' || char > '9' {
			return false
		}
	}
	return true
}
