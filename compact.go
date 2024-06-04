package piscine

// Function that takes a pointer to a slice of string and returns the
// number of elements with non-zero value and delete the elements with
// zero-values in the slice.
func Compact(ptr *[]string) int {
	count := 0
	for i := 0; i < len(*ptr); {
		if (*ptr)[i] != "" {
			count++
			i++
			// Take the slice until non-zero element and append the slice after zero element
		} else {
			*ptr = append((*ptr)[:i], (*ptr)[i+1:]...)
		}
	}
	return count
}
