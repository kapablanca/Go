package piscine

/* Function that counts the runes of a string
and returns that count */
func StrLen(s string) int {
	count := 0

	for range s {
		count++
	}
	return count
}
