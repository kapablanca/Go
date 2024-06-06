package piscine

// Function that takes a slice of strings(menu) and returns
// another slice with the items in the correct order.
// append() is not allowed for this exercise.
func ReverseMenuIndex(menu []string) []string {
	// Make a copy of the menu, so the changes will not affect our original slice
	newMenu := make([]string, len(menu))
	copy(newMenu, menu)
	// Change front and back items at the same time
	for i, j := 0, len(newMenu)-1; i < len(newMenu)/2; i, j = i+1, j-1 {
		newMenu[i], newMenu[j] = newMenu[j], newMenu[i]
	}

	return newMenu
}
