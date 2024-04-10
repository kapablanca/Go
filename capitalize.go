package piscine

func Capitalize(s string) string {
	var answer string
	var last_letter string
	for _, letter := range s {
		letter := string(letter)

		if !IsAlpha(letter) || IsNumeric(letter) || letter == " " {
			answer += letter
		} else {
			if answer == "" {
				if IsNumeric(letter) {
					answer += letter
				} else {
					answer += ToUpper(letter)
				}
			} else {
				last_letter = string(LastRune(answer))
				if IsNumeric(last_letter) || IsUpper(last_letter) || IsLower(last_letter) {
					answer += letter
				} else {
					answer += ToUpper(letter)
				}

			}
		}

	}
	return answer
}
