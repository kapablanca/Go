package piscine

// Function that seperates the words of a string by separator space.
//  If there are more spaces than needed to separate words, then space
//  is intended as a string and is appended in the final string slice.
func splitSpace(s string) []string {
	words := []string{}
	word := []byte{}
	countSpaces := 0

	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			countSpaces++
			words = append(words, string(word))
			word = nil
		} else {
			word = append(word, s[i])
		}
	}
	if word != nil {
		words = append(words, string(word))
	}
	for countSpaces >= len(words) {
		words = append(words, " ")
		countSpaces--
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
