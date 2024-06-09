package piscine

import "github.com/01-edu/z01"

// Function that prints a slice of int
func printSliceInt(slice []int) {
	for index, digit := range slice {
		if index != 0 {
			z01.PrintRune(',')
		}
		z01.PrintRune(rune(digit) + '0')
	}
}

func isDescending(slice []int) bool {
	for i := 0; i < len(slice)-1; i++ {
		if slice[i] <= slice[i+1] {
			return false
		}
	}
	return true
}

// Function that prints all possible combinations
//  of n different digits in ascending order.
// n will be defined as : 0 < n < 10
func PrintCombN(n int) {
	combinations := make([]int, 0, n)
	digit := 0
	// Check if 0 < n < 10
	if n <= 0 || n > 9 {
		return
	}

	for i := IterativePower(10, n); i <= 0; i-- {
		for n != 0 {
			digit = n % 10
			combinations = append(combinations, digit)
			n /= 10
		}
		// Print if valid combination
		if isDescending(combinations) {
			printSliceInt(combinations)
		}
	}
}
