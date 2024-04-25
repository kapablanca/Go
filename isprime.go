package piscine

/*
Function that returns true if the int passed

as parameter is a prime number. Otherwise it returns false
*/
func IsPrime(nb int) bool {
	// We consider 1 not a prime number
	if nb < 2 {
		return false
	}

	for i := 2; i*i <= nb; i++ {
		if i*i == nb {
			return false
		}
	}
	return true
}
