package piscine

func AppendRange(min, max int) []int {
	var series []int

	if min >= max {
		return series
	}
	for i := min; i < max; i++ {
		series = append(series, i)
	}
	return series
}
