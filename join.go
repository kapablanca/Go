package piscine

// Function that returns the concatenation of all the strings
// of a slice of string seperated by the sep
func Join(strs []string, sep string) string {
	joined := ""
	for index, str := range strs {
		if index == 0 {
			joined += str
		} else {
			joined += sep + str
		}
	}
	return joined
}
