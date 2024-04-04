package piscine

func StrLen(s string) int {
	aString := []byte(s)

	count := 0

	for range aString {
		count++
	}

	return count
}
