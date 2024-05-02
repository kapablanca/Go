package piscine

// Function that capitalizes the first letter of
// each word and lowercases the rest. A word is a
// sequence of alphanumeric characters.
func Capitalize(s string) string {
	capitalized := ""
	word := ""

	for _, char := range s {
		substring := string(char)
		// Check if the substring is alphanumeric char
		if IsAlpha(substring) {
			// Capitalize the first letter of word if needed
			if word == "" && IsLower(substring) {
				substring = ToUpper(substring)
			}
			// Lowercase the letter if needed
			if word != "" && IsUpper(substring) {
				substring = ToLower(substring)
			}
			word += substring
			// Append the non alpha char and clear the word variable
		} else {
			capitalized += word + substring
			word = ""
		}
	}
	capitalized += word
	return capitalized
}
