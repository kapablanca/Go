package piscine

import "fmt"

// Function that returns the element of the slice that does not have a
//  correspondet pair. If all the numbers have a correspondent pair, it should return -1
func Unmatch(a []int) int {
	// Create a map to count how many times a number appears
	numbersFreq := make(map[int]int)
	// Counting numbers
	for _, i := range a {
		numbersFreq[i]++
	}
	fmt.Print(numbersFreq)
	// Finding pairs
	for key, value := range numbersFreq {
		if value%2 != 0 {
			return key
		}
	}
	return -1
}
