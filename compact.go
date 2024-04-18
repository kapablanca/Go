package piscine

func Compact(ptr *[]string) int {
	count := 0

	for i := 0; i < len(*ptr); i++ {
		if (*ptr)[i] != "" {
			count++
		} else {
			(*ptr) = append((*ptr)[:i], (*ptr)[i+1:]...)
		}
	}
	return count
}
