package piscine

func JumpOver(str string) string {
	var result string
	if str == "" || len(str) < 3 {
		result = "\n"
	} else {
		for index, char := range str {
			if (index+1)%3 == 0 && index+1 < len(str) {
				result += string(char)
			}
		}
		result += "\n"
	}

	return result
}
