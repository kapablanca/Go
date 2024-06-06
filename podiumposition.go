package piscine

import "fmt"

// Function that takes an unorderd slice of slices  of string and returns
// the competitor's positions correctly.
func HasChar(s string, r rune) bool {
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
	// Make an emptry 2d array with same dimensions as podium
	newPodium := make([][]string, rows)
	for index := range newPodium {
		newPodium[index] = make([]string, cols)
	}
	//
	for range newPodium {
		for irow, slice := range newPodium {
			fmt.Println(irow)
			fmt.Println(rune(irow+1) + '0')
			fmt.Printf("Slice[0] is :%v", slice[0][irow])

			if HasChar(slice[0], rune(irow+1)+'0') {
				newPodium[irow][0] = slice[0]

			}

		}
	}
	return newPodium
}
