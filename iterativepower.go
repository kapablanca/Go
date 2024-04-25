package piscine

/*
	Function that returns the value of nb to the power of power.

Negative powers will return 0. Overflows do not have to be dealt with
*/
func IterativePower(nb int, power int) int {
	// Check if power is negative
	if power < 0 {
		return 0
	}

	number := 1
	// Number to the power of 0 is 1
	if power == 0 {
		return number
	}

	// Calculating  the result
	for i := 0; i < power; i++ {
		number *= nb
	}
	return number
}
