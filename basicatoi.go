package piscine

/* Function that transforms a number defined as string into an integer.
For this exercise I take into account only valid strings and don't take into
account the handling of signs */
func BasicAtoi(s string) int {
	number := 0
	digit := 0

	// Turning each rune into an int and adding it to the number
	for _, char := range s {
		digit = int(char - '0')
		number = 10*number + digit
	}
	return number
}
