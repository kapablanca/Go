package piscine

import "github.com/01-edu/z01"

func PrintNbrBase(nbr int, base string) {
	// Checking length of base
	base_len := len(base)
	if base_len < 2 {
		PrintNV()
		return
	}
	// Checking unique characters or sign inside base
	for _, character := range base {
		if !IsUnique(base, string(character)) || character == '+' || character == '-' {
			PrintNV()
			return
		}
	}

	// check if number is negative to add the "-" sign later
	negative := false

	if nbr < 0 {
		nbr = nbr * (-1)
		negative = true
	}

	// empty slice to save the digits from the transformation
	var digits_save []int
	var reminder int
	var quotient int = nbr

	// Tranforming the digits
	for quotient != 0 {
		reminder = quotient % base_len
		digits_save = append(digits_save, reminder)
		quotient /= base_len
	}

	// adding the minus sign in front of the number
	if negative {
		z01.PrintRune('-')
	}

	// printint the digits to the new base
	for i := len(digits_save) - 1; i >= 0; i-- {
		for j, character := range base {
			if digits_save[i] == j {
				z01.PrintRune(character)
			}
		}
	}
}

func IsUnique(word string, character string) bool {

	for i, char1 := range word {

		for j, char2 := range word {

			if i != j && char1 == char2 {
				return false
			}
		}
	}
	return true
}

func PrintNV() {

	z01.PrintRune('N')
	z01.PrintRune('V')
}
