package piscine

// Functioni that retunrs the number of elements of a string slice,
// for which the f function returns true.
func CountIf(f func(string) bool, tab []string) int {
	count := 0
	for _, elem := range tab {
		if f(elem) {
			count++
		}
	}
	return count
}
