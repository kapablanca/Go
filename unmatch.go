package piscine

func Unmatch(a []int) int {
	counts := make(map[int]int)

	for _, nbr := range a {
		counts[nbr]++
	}

	for _, nbr := range a {
		if counts[nbr]%2 != 0 {
			return nbr
		}
	}
	return -1
}
