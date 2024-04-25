package piscine

/*
Iterative function that returns the factorial
of the int passed as parameter. Errors will return 0
*/
func IterativeFactorial(nb int) int {
	factorial := 1
	// Check if nb is positive
	if nb < 0 {
		return 0
	}
	// Zero factorial is equal to 1
	if nb == 0 {
		return factorial
	}
	// Calculate factorial
	for i := 1; i <= nb; i++ {
		factorial *= i
		// Check for overflow
		if factorial < 0 {
			return 0
		}
	}
	return factorial
}
