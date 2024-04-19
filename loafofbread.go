package piscine

func LoafOfBread(str string) string {
	if len(str) < 5 {
		return "Invalid Output\n"
	}

	var result string
	count := 0
	groupCharCount := 0

	for i, char := range str {
		if char != ' ' {
			count++
			if count <= 5 {
				result += string(char)
				groupCharCount++
			} else if count == 6 {
				count = 0
				if i+1 < len(str) {
					result += " "
				}
			}
		}
	}
	return result + "\n"
}
