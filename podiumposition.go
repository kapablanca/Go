package piscine

// Function that outputs if a string has a specific rune
func hasChar(s string, r rune) bool {
	for _, char := range s {
		if char == r {
			return true
		}
	}
	return false
}

// Function that takes an unorderd slice of slices  of string and returns
// the competitor's positions correctly.
func PodiumPosition(podium [][]string) [][]string {
	// rows := len(podium)
	// cols := len(podium[0])
	// Slice that holds the podium positions
	positions := []rune{'1', '2', '3', '4'}
	// Make an emptry 2d array with same dimensions as podium
	// var newPodium [][]string
	// for i := range podium {
	// 	newPodium[i] = []string{""}
	// }
	// Search the position on each string and put it in the correct order
	for index, char := range positions {
		for rowIndex, slice := range podium {
			for _, str := range slice {
				if hasChar(str, char) {
					// newPodium[index][0] = str
					podium[index][0], podium[rowIndex][0] = podium[rowIndex][0], podium[index][0]
				}
			}
		}
	}
	return podium
}
