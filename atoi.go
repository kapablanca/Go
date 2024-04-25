package piscine

/* Function that simulates the behaviour of the Atoi in Go.
Atoi transforms a number represented as a string in a number
represented as an int. Atoi returns 0 if the string is not considered
a valid number */

// Checks if a rune char is number or not
func isNumber(nbr rune) bool {
	if nbr < '0' || nbr > '9' {
		return false
	}
	return true
}

func Atoi(s string) int {
	number := 0
	digit := 0
	sign := 1

	// Checking if the string is empty
	if s == "" {
		return 0
	}

	// Checking if there is a sign at the beginning
	if len(s) > 1 {

		if (s[0] == '-' || s[0] == '+') && isNumber(rune(s[1])) {
			sign = -1
			// Removing it from the string
			s = s[1:]
		}
	}

	for _, char := range s {
		// Checking if the char represents a valid number
		if isNumber(char) {
			digit = int(char - '0')
			number = 10*number + digit
			// Return 0 if not valid char
		} else {
			return 0
		}
	}
	return sign * number
}
