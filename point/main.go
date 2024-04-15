package main

import "github.com/01-edu/z01"

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func numberToString(nbr int) string {
	var digits []rune
	var digitsString string

	mod := 1

	for mod != 0 {
		mod = nbr % 10
		if mod == 0 {
			continue
		}
		digits = append(digits, rune(mod)+48)
		nbr /= 10
	}

	for i := len(digits) - 1; i >= 0; i-- {
		digitsString += string(digits[i])
	}
	return digitsString
}

func main() {
	points := point{}
	points_point := &points

	setPoint(points_point)

	x := numberToString(points.x)
	y := numberToString(points.y)

	sentence := "x = " + x + ", y = " + y + "\n"

	for _, char := range sentence {
		z01.PrintRune(char)
	}
}
