package piscine

func IsPrime(nb int) bool {
	result := true
	for i := 2; i < nb; i++ {
		if nb%i == 0 {
			result = false
		}
	}
	return result
}
