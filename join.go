package piscine

func Join(strs []string, sep string) string {
	var answer string

	for _, word := range strs {
		if answer == "" {
			answer += word
		} else {
			answer += sep + word
		}
	}
	return answer
}
