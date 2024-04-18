package piscine

func SplitOneSpace(str string) []string {
	var words []string
	currentWord := ""
	inWord := false

	for i := 0; i < len(str); i++ {
		if str[i] != ' ' {
			currentWord += string(str[i])
			inWord = true
		} else if inWord {
			// Only add the currentWord to words if we were previously in a word.
			words = append(words, currentWord)
			currentWord = ""
			inWord = false
		}
	}

	// Add the last word if it exists
	if inWord {
		words = append(words, currentWord)
	}

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
