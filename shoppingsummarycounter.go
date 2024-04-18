package piscine

func SplitOneSpace(str string) []string {
	var words []string

	for i := 0; i < len(str); {
		if string(str[i]) == " " {

			words = append(words, str[:i])
			str = str[i+1:]
			i = 0
		} else {
			i++
		}
	}
	if len(str) > 0 {
		words = append(words, str)
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
