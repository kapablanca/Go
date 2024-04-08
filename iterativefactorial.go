package piscine

func IterativeFactorial(nb int) int {
	if nb < 0 {
		return 0
	}
	var result int = 1
	var maxInteger int = 2147483647

	if nb == 0 {
		return result
	}

	for i := 1; i <= nb; i++ {
		if result*i > maxInteger {
			return 0
		} else {
			result = result * i
		}
	}
	return result
}
