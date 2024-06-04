package piscine

// Function that returns the median of five int arguments
func Abort(a, b, c, d, e int) int {
	numbers := []int{a, b, c, d, e}
	// Sort the values to find the median
	SortIntegerTable(numbers)
	// Median is the middle value(3rd value in 5 elements)
	return numbers[2]
}
