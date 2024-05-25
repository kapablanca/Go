package piscine

// Function that returns true for a string slice, if
//  when that string slice is passed through an f function,
// at least one elements returns true

func Any(f func(string) bool, a []string) bool {
	for _, elem := range a {
		if f(elem) {
			return true
		}
	}
	return false
}
