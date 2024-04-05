package piscine

func BasicAtoi2(s string) int {
	var number int = 0

	for _, num := range s {

		var digit rune = num - 48
		if int(digit) <= 0 && int(digit) >= 9 {
			return 0
		}
		number = number*10 + int(digit)
	}

	return number
}
