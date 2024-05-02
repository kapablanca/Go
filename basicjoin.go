package piscine

// Function that returns a concatenated string from the
//  strings passed in as arguments
func BasicJoin(elems []string) string {
	concatenated := ""
	for _, elem := range elems {
		concatenated += elem
	}
	return concatenated
}
