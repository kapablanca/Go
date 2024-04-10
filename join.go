package piscine

func Join(strs []string, sep string) string {
	var answer string

	for _, word := range strs {
		answer += word + sep
	}

	if string(LastRune(answer)) == sep {
		answer = answer[:len(answer)-1]
	}
	return answer
}
