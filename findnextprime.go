package piscine

/* Function that returns the first prime number that is equal or
superior to the int passed as parameter. */
func FindNextPrime(nb int) int {

	if IsPrime(nb) {
		return nb
	}
	// Recursively check the next number
	return FindNextPrime(nb + 1)
}
