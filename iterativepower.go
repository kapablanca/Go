package piscine

func IterativePower(nb int, power int) int {
	if power < 0 {
		return 0
	}
	if nb == 0 {
		return 1
	}
	result := 1
	for i := 1; i < power+1; i++ {
		result *= nb
	}
	return result
}
