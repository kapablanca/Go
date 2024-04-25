package piscine

/*
Function that returns the factorial of the int
passed as parameter. Errors will return 0
*/
func RecursiveFactorial(nb int) int {
	// Check if negative number
	if nb < 0 {
		return 0
	}
	// Factorial of zero is 1
	if nb == 0 {
		return 1
	}
	// Return the number times the factorial of the previous number
	// n! = n * (n-1)!
	return nb * RecursiveFactorial(nb-1)
}
