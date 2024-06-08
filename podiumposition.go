package piscine

// Return the first char of a string
func firstChar(s string) byte {
	if len(s) == 0 {
		return 0
	}
	return s[0]
}

// Function that takes an unorderd slice of slices  of string and returns
// the competitor's positions correctly.
func PodiumPosition(podium [][]string) [][]string {
	// Sort the 2d array based on first value of string that shows podium position
	for range podium {
		for i := 0; i < len(podium)-1; i++ {
			if firstChar(podium[i][0]) > firstChar(podium[i+1][0]) {
				podium[i][0], podium[i+1][0] = podium[i+1][0], podium[i][0]
			}
		}
	}
	return podium
}
