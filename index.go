package piscine

// Function that returns the index of the first instance
// of toFind in s, or -1 if toFind not present in s
func Index(s string, toFind string) int {
	n := len(toFind)
	for i := 0; i < len(s) - n; i++ {
		if s[i:i + n] == toFind {
			return i
		}
	}
	return -1
}