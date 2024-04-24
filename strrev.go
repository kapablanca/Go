package piscine

/* Function that reverses a string
and returns it */
func StrRev(s string) string {
	reversedString := ""
	sLen := StrLen(s)

	for i := sLen - 1; i >= 0; i-- {
		reversedString += string(rune(s[i]))
	}
	return reversedString
}
