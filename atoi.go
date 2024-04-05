package piscine

func Atoi(s string) int {
	if s == "" {
		return 0
	}

	number := 0
	sign := 1
	start := 0

	// Check if the string starts with a sign.
	if s[0] == '-' || s[0] == '+' {
		if s[0] == '-' {
			sign = -1
		}
		start = 1
	}

	for i := start; i < len(s); i++ {
		digit := s[i] - '0'
		if digit < 0 || digit > 9 {
			return 0 // Contains a non-digit character.
		}
		number = number*10 + int(digit)
	}

	return number * sign
}
