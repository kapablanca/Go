package piscine

func Atoi(s string) int {
	var number int = 0
	sign := s[0]
	if (sign == '-' || sign == '+') && s[1] != '+' || s[1] != '-' {
		s = s[1:]
	} else {
		return 0
	}

	for _, num := range s {

		var digit rune = num - 48
		if int(digit) < 0 || int(digit) > 9 {
			return 0
		}
		number = number*10 + int(digit)
	}

	if sign == '-' {
		return number * (-1)
	} else {
		return number
	}
}