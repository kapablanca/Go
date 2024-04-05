package piscine

func BasicAtoi(s string) int {
	var number int = 0

	for _, num := range s {

		var digit rune = num - 48
		number = number*10 + int(digit)
	}

	return number
}
