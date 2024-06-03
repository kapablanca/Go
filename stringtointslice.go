package piscine

// Function that takes a string and returns the corresponding int slice
func StringToIntSlice(str string) []int {
	intSlice := []int{}

	for _, char := range str {
		intSlice = append(intSlice, int(char))
	}
	return intSlice
}
