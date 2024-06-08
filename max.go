package piscine

// Function that returns the maximum value in a slice of integers.
// If the slice is empty, return 0.
func Max(a []int) int {
	// Check if slice is empty
	if len(a) == 0 {
		return 0
	}
	max := 0
	// Compare each integer to the var max
	for _, value := range a {
		if value > max {
			max = value
		}
	}
	return max
}
