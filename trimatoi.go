package piscine

// Function that transforms numbers within a string into an int.
// If the - sign is encountered before any number, it should determine
// the sign of the returned int. In case of invalid input, the function
// should return 0.
func TrimAtoi(s string) int {
	number := 0
	isValid := false
	sign := 1

	for _, char := range s {
		// negative sign if it hasn't found a number yet
		if char == '-' && !isValid {
			sign = -1
		}
		if IsNumeric(string(char)) {
			isValid = true
			number = number*10 + int(char-'0')
		}
	}

	return sign * number
}
