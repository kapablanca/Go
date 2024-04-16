package piscine

func Map(f func(int) bool, a []int) []bool {
	var listBools []bool

	for _, nbr := range a {
		listBools = append(listBools, f(nbr))
	}

	return listBools
}
