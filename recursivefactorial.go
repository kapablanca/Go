package piscine

/*
Function that returns the factorial of the int
passed as parameter. Errors will return 0
*/
func RecursiveFactorial(nb int) int {
	var factorial int
	// Check if negative number
	if nb < 0 {
		return 0
	}
	// Factorial of zero is 1
	if nb == 0 {
		factorial = 1
	}
	factorial *= RecursiveFactorial(nb - 1)
	return factorial
}
