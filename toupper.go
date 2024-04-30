package piscine

// Function that capitalizes each letter in a string
func ToUpper(s string) string {
	chars := []rune(s)

	for i := 0; i < len(chars); i++ {
		if IsLower(string(chars[i])) {
			chars[i] -= ' '
		}
	}
	return string(chars)
}
