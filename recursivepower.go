package piscine

/*
Function that recursively calculates the value of nb to

the power of power. Negative powers will return 0. Overflows

do not have to be dealt with.
*/
func RecursivePower(nb int, power int) int {
	// Checking if the power is negative
	if power < 0 {
		return 0
	}

	if power == 0 {
		return 1
	}
	return nb * RecursivePower(nb, power-1)
}
