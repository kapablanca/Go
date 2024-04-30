package piscine

// Function that returns true if the string passed as a parameter
// contains only printable characters, else returns false
// (Non printable characters ASCII 0-31 and 127)
func IsPrintable(s string) bool {
	for _, char := range s {
		if char <= 31 || char == 127 {
			return false
		}
	}
	return true
}
