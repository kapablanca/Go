package piscine

func AtoiBase(s string, base string) int {
	// Checking length of base
	base_len := len(base)
	if base_len < 2 {
		return 0
	}
	// Checking unique characters or sign inside base
	for _, character := range base {
		if !IsUnique(base, string(character)) || character == '+' || character == '-' {
			return 0
		}
	}

	// Stating variables
	var numbers []int
	var index_numerical_value int
	var number int
	var power int
	var integer_number int

	// Checking
	for i, digit := range s {

		power = len(s) - 1 - i

		for j, digit_base := range base {
			if digit == digit_base {

				index_numerical_value = j
				number = index_numerical_value * IterativePower(len(base), power)
				numbers = append(numbers, number)

			}
		}
	}
	for _, number := range numbers {
		integer_number += number
	}

	return integer_number
}
