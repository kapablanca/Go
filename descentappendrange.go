package piscine

func DescendAppendRange(max, min int) []int {
	var list []int

	if max <= min {
		return list
	}
	for i := max; i > min; i-- {
		list = append(list, i)
	}
	return list
}
