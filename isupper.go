package piscine

// Function that returns true if the string passed as
// parameter contains only uppercase characters, else false
func IsUpper(s string) bool {
	for _, char := range s {
		if char < 'A' || char > 'Z' {
			return false
		}
	}
	return true
}
