package piscine

func StringToIntSlice(str string) []int {
	var list []int

	for _, char := range str {
		list = append(list, int(char))
	}

	return list
}
