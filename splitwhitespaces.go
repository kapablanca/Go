package piscine

func SplitWhiteSpaces(s string) []string {
	var splitted []string
	var word string

	for i := 0; i < len(s); i++ {
		if s[i] == byte(rune(32)) || s[i] == byte(rune(32)) || s[i] == byte(rune(32)) {
			if i != 0 && (s[i-1] != byte(rune(32)) || s[i-1] != byte(rune(32)) || s[i-1] != byte(rune(32))) {

				splitted = append(splitted, word)
				word = ""
			} else {
				word = ""
			}
		} else {

			word += string(s[i])

			if i == len(s)-1 {
				splitted = append(splitted, word)
				word = ""
			}

		}
	}

	return splitted
}
