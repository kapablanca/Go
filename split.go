package piscine

// Function that receives a string and a separator and returns a slice of strings
// that results of splitting the string s by the separator sep
func Split(s, sep string) []string {
	var words []string
	var word []rune
	// Extending the string to stop the iteration in the end
	s += sep

	for i := 0; i <= len(s)-len(sep); i++ {
		if s[i:i+len(sep)] == sep {
			// Make the next iteration pass the seperator by changing index
			i += len(sep) - 1
			// Go to the next iter if there is no word to append
			if word == nil {
				continue
			}
			// Extract the word and append it to the slice words
			words = append(words, string(word))
			word = nil
		} else {
			word = append(word, rune(s[i]))
		}
	}
	return words
}
