package piscine

func Unmatch(a []int) int {
	unique := 0

	for i := 0; i < len(a); i++ {
		unique = unique ^ a[i]
	}
	if unique == 0 {
		return -1
	}
	return unique
}
