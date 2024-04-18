package piscine

func Unmatch(a []int) int {
	counts := make(map[int]int)

	for _, nbr := range a {
		counts[nbr]++
	}

	for key, count := range counts {
		if count%2 != 0 {
			return key
		}
	}
	return -1
}
