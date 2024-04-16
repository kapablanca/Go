package piscine

func f(a, b int) int {
	if a == b {
		return 0
	} else if a > b {
		return 1
	} else {
		return -1
	}
}

func IsSorted(f func(a, b int) int, a []int) bool {
	var sorting_list []int
	one := 0
	negative_one := 0

	for i := 0; i < len(a)-1; i++ {
		// for j := i + 1; j < len(a); j++ {
		// 	sorting_list = append(sorting_list, f(a[i], a[j]))
		// }
		sorting_list = append(sorting_list, f(a[i], a[i+1]))
	}

	for _, digit := range sorting_list {
		if digit == 1 {
			one = 1
		}
		if digit == -1 {
			negative_one = 1
		}
	}
	if one == 1 && negative_one == 1 {
		return false
	}
	return true
}
