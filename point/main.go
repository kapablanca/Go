package main

import (
	"github.com/01-edu/z01"
)

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func nbToString(nb int) string {
	var answer []rune
	var mod int
	var digit rune

	for nb != 0 {
		mod = nb % 10
		digit = rune(mod) + '0'
		answer = append(answer, digit)
		nb /= 10
	}
	// Reverse the digits of slice
	for i, j := 0, len(answer)-1; i < len(answer)/2; i, j = i+1, j-1 {
		answer[i], answer[j] = answer[j], answer[i]
	}
	return string(answer)
}
func printPoint(x, y int) {
	xString := nbToString(x)
	yString := nbToString(y)

	xMessage := "x = " + xString
	yMessage := "y = " + yString
	message := xMessage + ", " + yMessage

	for _, char := range message {
		z01.PrintRune(char)
	}
	z01.PrintRune('\n')
}

// The necessary changes must be applied so that the program works.
// The function setPoint() must work with int.
func main() {
	points := &point{}

	setPoint(points)

	printPoint(points.x, points.y)
}
