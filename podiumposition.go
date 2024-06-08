package piscine

// Function that takes an unorderd slice of slices  of string and returns
// the competitor's positions correctly.
func hasChar(s string, r rune) bool {
	for _, char := range s {
		if char == r {
			return true
		}
	}
	return false
}
func PodiumPosition(podium [][]string) [][]string {
	rows := len(podium)
	cols := len(podium[0])
	// Slice that holds the podium positions
	positions := []rune{'1', '2', '3', '4'}
	// Make an emptry 2d array with same dimensions as podium
	newPodium := make([][]string, rows)
	for i := range podium {
		newPodium[i] = make([]string, cols)
	}
	// Search the position on each string and put it in the correct order
	for index, char := range positions {
		for _, slice := range podium {
			for _, str := range slice {
				if hasChar(str, char) {
					newPodium[index][0] = str
				}
			}
		}
	}
	return newPodium
}
