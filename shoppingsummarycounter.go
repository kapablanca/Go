package piscine

func ShoppingSummaryCounter(str string) map[string]int {
	items_slice := SplitWhiteSpaces(str)

	shopping_list := make(map[string]int)

	for _, item := range items_slice {
		shopping_list[item]++
	}

	return shopping_list
}
