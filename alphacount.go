package piscine

// Function that returns true if the rune is a letter
func isLetter(char rune) bool {
	if (char >= 'A' && char <= 'Z') ||
		(char >= 'a' && char <= 'z') {
		return true
	}
	return false
}

// Function that counts the lettes of a string
// and returns the count
func AlphaCount(s string) int {
	count := 0
	for _, char := range s {
		if isLetter(char) {
			count++
		}
	}
	return count
}
