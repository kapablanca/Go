package piscine

// Function that returns the element of the slice that does not have a
//  correspondet pair. If all the numbers have a correspondent pair, it should return -1
func Unmatch(a []int) int {
	// Create a map to count how many times a number appears
	numbersFreq := make(map[int]int)
	// Counting numbers
	for _, i := range a {
		numbersFreq[i]++
	}
	// Finding if the number is paired by iterating in the slice order
	for _, i := range a {
		if numbersFreq[i]%2 != 0 {
			return i
		}
	}
	return -1
}
