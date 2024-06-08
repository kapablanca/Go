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
	var position rune
	// Search the position on each string and put it in the correct order
	for index := range podium {
		// Make the index a podium position
		position = rune(index+1) + '0'
		for rowIndex, slice := range podium {
			for _, str := range slice {
				if hasChar(str, position) {
					podium[index][0], podium[rowIndex][0] = podium[rowIndex][0], podium[index][0]
				}
			}
		}
	}
	return podium
}
