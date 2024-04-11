package piscine

func MakeRange(min, max int) []int {
	var series []int

	if min >= max {
		return series
	} else {
		size := max - min
		series = make([]int, size)

		for i := 0; i < size; i++ {
			series[i] = i + min
		}
	}
	return series
}
