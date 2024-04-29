package piscine

/*
Function that returns an integer comparing
*/
func Compare(a, b string) int {
	if a == b {
		return 0
	} else if a > b {
		return 1
	}
	return -1
}
