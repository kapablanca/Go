package piscine

// Function that returns the concatenation of all the strings
// of a slice of string seperated by the sep
func Join(strs []string, sep string) string {
	joined := ""

	for index, word := range strs {
		// Add separator before every concat
		if index != 0 {
			joined += sep
		}
		joined += word
	}
	return joined
}
