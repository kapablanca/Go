package piscine

// Function that lower cases for each letter in a string
func ToLower(s string) string {
	chars := []rune(s)

	for i := 0; i < len(chars); i++ {
		if IsUpper(string(chars[i])) {
			chars[i] += ' '
		}
	}
	return string(chars)
}
