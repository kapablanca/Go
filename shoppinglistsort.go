package piscine

// Function that takes a slice of strings and sorts it,
// according to the string length, in ascending order.
// Strings within the inpyt slice must be of different lengths.
func ShoppingListSort(slice []string) []string {
	// Use bubble sort to the slice
	for range slice {
		for i := 0; i < len(slice)-1; i++ {
			if len(slice[i]) > len(slice[i+1]) {
				slice[i], slice[i+1] = slice[i+1], slice[i]
			}
		}
	}
	return slice
}
