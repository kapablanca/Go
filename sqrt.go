package piscine

/*
Function that returns the square root of the int passed as parameter,

if that square root is a whole number. Otherwise it returns 0
*/
func Sqrt(nb int) int {
	// Check if number is greater than zero
	if nb <= 0 {
		return 0
	}
	// Finding the sqaure root by checking if a number
	// squared is equal to the nb
	for root := 1; root*root <= nb; root++ {
		if root*root == nb {
			return root
		}
	}
	return 0
}
