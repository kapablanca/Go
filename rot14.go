package piscine

func Rot14(s string) string {
	var answer string

	for _, char := range s {
		if IsUpper(string(char)) {
			temp := rune(char) + 14
			if temp > rune('Z') {
				temp -= 26
			}
			answer += string(temp)

		} else if IsLower(string(char)) {
			temp := rune(char) + 14
			if temp > rune('z') {
				temp -= 26
			}
			answer += string(temp)
		} else {
			answer += string(char)
		}
	}
	return answer
}
