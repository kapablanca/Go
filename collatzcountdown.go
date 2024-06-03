package piscine

// Function that returns the number of steps necessary to reach 1
// using the collatz countdown. It must return -1 if start is
// equal to 0 or negative.
func CollatzCountdown(start int) int {
	steps := 0
	// Check if start is non positive
	if start <= 0 {
		return -1
	}
	if start == 1 {
		steps++
		return steps
	}
	for start != 1 {
		// Start is even
		if start%2 == 0 {
			start /= 2
			// Start is odd
		} else {
			start = 3*start + 1
		}
		steps++
	}
	return steps
}
