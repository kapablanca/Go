package piscine

// Function that takes a slice of strings(menu) and returns
// another slice with the items in the correct order.
// append() is not allowed for this exercise.
func ReverseMenuIndex(menu []string) []string {
	newMenu := make([]string, len(menu))
	// Change front and back items at the same time
	for i, j := 0, len(menu)-1; i < len(menu)/2; i, j = i+1, j-1 {
		menu[i], menu[j] = menu[j], menu[i]
	}
	copy(newMenu, menu)
	return newMenu
}
