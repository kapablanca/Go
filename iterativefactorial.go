package piscine

func IterativeFactorial(nb int) int {
	if nb < 0 {
		return 0
	}
	var result int = 1

	for i := 1; i <= nb; i++ {
		if nb < 21 {
			result = result * i
		} else {
			return 0
		}
	}
	return result
}
