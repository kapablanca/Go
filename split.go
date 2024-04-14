package piscine

func Split(s, sep string) []string {
	var words []string

	for i := 0; i < len(s)-len(sep); {

		j := len(sep)

		if s[i:i+j] == sep {

			if i != 0 {
				words = append(words, s[:i])
			}
			s = s[i+j:]
			i = 0
		} else {
			i++
		}

	}
	if len(s) > 0 {
		words = append(words, s)
	}
	return words
}
