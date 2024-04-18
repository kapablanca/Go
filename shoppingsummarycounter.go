package piscine

func SplitOneSpace(str string) []string {
	var words []string
	currentWord := ""

	if str == " " {
		words = append(words, "")
	}

	// Process each character in the string
	for i := 0; i < len(str); i++ {
		if str[i] == ' ' {
			// Check if currentWord has accumulated any characters before adding it to words
			// if currentWord != "" {
			words = append(words, currentWord)
			currentWord = "" // Reset currentWord after adding it to words
			// }
			// Continue to next iteration to skip additional space processing
			continue
		}
		// Add non-space characters to currentWord
		currentWord += string(str[i])
	}

	// Add the last accumulated word to words if not empty

	words = append(words, currentWord)

	return words
}

func ShoppingSummaryCounter(str string) map[string]int {
	items_slice := SplitOneSpace(str)

	shopping_list := make(map[string]int)

	for _, item := range items_slice {
		shopping_list[item]++
	}

	return shopping_list
}
