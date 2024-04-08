package piscine

func IterativeFactorial(nb int) int {
	if nb < 0 {
		return 0
	}
	var result int = 1
	var maxInteger int = 9223372036854775807

	for i := 1; i <= nb; i++ {
		if result < maxInteger {
			result = result * i
		} else {
			return 0
		}
	}
	return result
}
