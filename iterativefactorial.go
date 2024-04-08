package piscine

func IterativeFactorial(nb int) int {
	if nb < 0 {
		return 0
	}
	var result int = 1
	var maxInteger int = 92147483647

	for i := 1; i <= nb; i++ {
		if result > maxInteger {
			return 0
		} else {
			result = result * i
		}
	}
	return result
}
