package piscine

// Function that counts the total amount of times each item appears in a string
// and returns a summary including the item and its number of appearances as an int
func ShoppingSummaryCounter(str string) map[string]int {
	shoppingList := Split(str, " ")
	summary := make(map[string]int)
	// Count the times an item appears on the list
	for _, item := range shoppingList {
		summary[item]++
	}
	return summary
}
