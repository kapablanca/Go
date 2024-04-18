package piscine

func SplitOneSpace(str string) []string {
	var words []string
	currentWord := ""

	for i := 0; i < len(str); i++ {
		if str[i] != ' ' {
			currentWord += string(str[i])
		} else if len(currentWord) > 0 {
			words = append(words, currentWord)
			currentWord = ""
		}
	}

	// Add the last word if it's not empty
	if len(currentWord) > 0 {
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
