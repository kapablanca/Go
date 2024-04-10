package piscine

func BasicJoin(elems []string) string {
	var answer string
	for _, word := range elems {
		answer += word
	}
	return answer
}
