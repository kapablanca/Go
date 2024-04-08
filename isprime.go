package piscine

func IsPrime(nb int) bool {
	result := true
	if nb <= 1 {
		result = false
		return result
	}
	for i := 2; i < nb; i++ {
		if nb%i == 0 {
			result = false
			return result
		}
	}
	return result
}
