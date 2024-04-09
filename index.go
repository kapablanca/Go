package piscine

func Index(s string, toFind string) int {
	s_slice := []rune(s)
	s_toFInd := []rune(toFind)
	length := len(s_toFInd)

	for i := 0; i+length < len(s_slice); i++ {
		if Compare(string(s_slice[i:i+length]), toFind) == 0 {
			return i
		}
	}

	return -1
}
