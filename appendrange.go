package piscine

// Function that takes an inmin and and an int max as parameters.
// The function should return a slice of int s with all the values
// between the min and max. Min is included, max is excluded.
// If min >= max, a nil slice is returned. Make is not allowed.
func AppendRange(min, max int) []int {
	var s []int
	if min >= max {
		return s
	}
	for i := min; i < max; i++ {
		s = append(s, i)
	}
	return s
}
