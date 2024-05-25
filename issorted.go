package piscine

// Function that returns true, if the slice of int is sorted, otherwise false.
// The function passed as an argument func(a, b int) returns a positive int if the
// first argument is greater than the second argument, it returns 0 if they are equal
// and it returns a negative int otherwise.
func IsSorted(f func(a, b int) int, a []int) bool {
	isPositive := false
	isNegative := false
	// Finding if there are a>b or b>a relations
	for i := 0; i < len(a)-1; i++ {
		if f(a[i], a[i+1]) > 0 {
			isPositive = true
		}
		if f(a[i], a[i+1]) < 0 {
			isNegative = true
		}
	}
	// If the list is both ascending and descending
	if isPositive && isNegative {
		return false
	}
	return true
}
