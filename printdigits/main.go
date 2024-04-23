package main

import "github.com/01-edu/z01"

// Print all the digits in ascending order
func main() {
	digits := "0123456789"

	for _, digit := range digits {
		z01.PrintRune(digit)
	}
	z01.PrintRune('\n')
}
