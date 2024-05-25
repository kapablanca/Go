package piscine

// Function that for an int slice applies a function on each element of that
// slice and returns a slice of all the return values
func Map(f func(int) bool, a []int) []bool {
	var slice []bool

	for _, elem := range a {
		slice = append(slice, f(elem))
	}
	return slice
}
