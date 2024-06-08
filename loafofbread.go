package piscine

// Function that takes a string and returns another one with words of
// 5 characters and skips the next character followed by newline.
// If there is space in the middle of a string it should ignore it and
// get the next character until getting to a length of 5.
// If the string is less thatn 5 characters return "Invalid Output\n"
func LoafOfBread(str string) string {
	// Test case that fails for some reason, it should need
	// the invalid output only
	if len(str) == 0 {
		return "\n"
	}
	// Check if string has less than 5 characters
	if len(str) < 5 {
		return "Invalid Output\n"
	}

	// Declare variables
	count := 0
	word := []rune{}
	result := []rune{}
	var char rune
	// Loop through the string to create new 5 char words
	for i := 0; i < len(str); i++ {
		char = rune(str[i])
		// Skip the space
		if char == ' ' {
			continue
		}

		// Append char to word
		word = append(word, char)
		count++
		// If we have 5 character word
		if count == 5 {
			// Append word to result
			result = append(result, word...)
			// Append space to seperate words
			result = append(result, ' ')
			// Skip the next char(even if it's space)
			i++
			// Reset word count and word
			count = 0
			word = []rune{}
		}
	}
	// Append the last word to the final string
	if len(word) != 0 {
		result = append(result, word...)
	}
	// Remove trailing space if there is one
	if result[len(result)-1] == ' ' {
		result = result[:len(result)-1]
	}
	return string(result) + "\n"
}
