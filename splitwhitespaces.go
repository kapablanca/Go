package piscine

// Function that separates the words of a string and puts them in
// a string slice. The separators are spaces, tabs and newlines.
func SplitWhiteSpaces(s string) []string {
	word := []rune{}
	words := []string{}

	for _, char := range s {
		if char == '\n' || char == '\t' || char == ' ' {
			// Go to the next iter if there is no word to append
			if word == nil {
				continue
			}
			// Extract the word and append it to the slice words
			words = append(words, string(word))
			word = nil
		} else {
			word = append(word, char)
		}
	}
	// Append the last word if not empty
	if word != nil {
		words = append(words, string(word))
	}
	return words
}
