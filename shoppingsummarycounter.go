package piscine

// Function that seperates the words of a string by separator space.
//  If there are more spaces than needed to separate words, then space
//  is intended as an empty string and is appended in the final string slice.
func splitSpace(s string) []string {
	words := []string{}
	word := []rune{}
	spaceBefore := false

	for _, char := range s {
		// Check if there is an empty string
		if char == ' ' && spaceBefore {
			words = append(words, string(word))
			spaceBefore = false
			// Append the word to words if there is a separator
		} else if char == ' ' {
			spaceBefore = true
			words = append(words, string(word))
			word = nil
			// Append char to word
		} else {
			spaceBefore = false
			word = append(word, char)
		}
	}
	// Append the last word if there is one to words
	if word != nil {
		words = append(words, string(word))
	}
	return words
}

// Function that counts the total amount of times each item appears in a string
// and returns a summary including the item and its number of appearances as an int
func ShoppingSummaryCounter(str string) map[string]int {
	shoppingList := splitSpace(str)
	summary := make(map[string]int)
	// Count the times an item appears on the list
	for _, item := range shoppingList {
		summary[item]++
	}
	return summary
}
