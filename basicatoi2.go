package piscine

/* Function that transforms a number defined as string into an integer.
BasicAtoi2 returns 0 if the string is not considered as a valid number
and don't take into account the handling of signs */
func BasicAtoi2(s string) int {
	number := 0
	digit := 0

	for _, char := range s {
		// Checking if the rune represents a valid number
		if char < '0' || char > '9' {
			return 0
		}

		digit = int(char - '0')
		number = 10*number + digit
	}
	return number
}
