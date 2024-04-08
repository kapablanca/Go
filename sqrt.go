package piscine

func Sqrt(nb int) int {
	if nb <= 0 {
		return 0
	}
	result := 0
	for i := 1; i*i <= nb; i++ {
		if i*i == nb {
			result = i
		}
	}
	return result
}
