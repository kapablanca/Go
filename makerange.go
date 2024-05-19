package piscine

// Function that takes an int min and int max as parameters.
// The function must return a slice of int s with all the values
// between min and max. Min is included, max is excluded.
// If min >= max, a nil slice is returned. Append is not allowed for this exercise.
func MakeRange(min, max int) []int {
	if min >= max {
		var s []int
		return s
	}
	s := make([]int, max-min)
	for index := range s {
		s[index] = min
		min++
	}
	return s
}
