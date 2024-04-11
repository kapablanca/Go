package piscine

func SplitWhiteSpaces(s string) []string {
	var splitted []string
	var word string

	for i, char := range s {
		if char == rune(32) || char == rune(10) || char == rune(9) {
			splitted = append(splitted, word)
			word = ""
		} else {

			word += string(char)

			if i == len(s)-1 {
				splitted = append(splitted, word)
				word = ""
			}

		}
	}

	return splitted
}
