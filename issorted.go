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
	for i := 0; i < len(a)-1; i++ {
		for j := i + 1; j < len(a); j++ {
			if f(a[i], a[j]) > 0 {
				return false
			}
		}
	}
	return true
}
