package piscine

func ToLower(s string) string {
	var answer string
	for _, letter := range s {
		if IsUpper(string(letter)) {
			answer += string(rune(letter + 32))
		} else {
			answer += string(letter)
		}
	}
	return answer
}
