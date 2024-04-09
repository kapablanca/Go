package piscine

func TrimAtoi(s string) int {
	num := 0
	sign := 1

	for _, char := range s {
		if char == '-' && num == 0 {
			sign = -1
		} else if char >= '0' && char <= '9' {
			num = num*10 + int(char-'0')
		}
	}
	return num * sign
}
