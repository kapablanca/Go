package piscine

// Function that returns a slice of int with all the values between
//  the max and min. The max must be included and min must be excluded.
// if max is inferior than or equal to min, return an empty slice.
// make() is not allowed for this exercise
func DescendAppendRange(max, min int) []int {
	list := []int{}
	// Check min, max
	if max <= min {
		return list
	}
	// Find the values between min-max
	for i := max; i > min; i-- {
		list = append(list, i)
	}
	return list
}
