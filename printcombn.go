package piscine

// Function that prints a slice of int
func generateComb(n int, digit rune, comb []rune) {
	if n == 0 {

	}
	for n > 0 {
		comb = append(comb, digit)
		generateComb(n-1, digit+1, comb)
	}
}

// Function that prints all possible combinations
//
//	of n different digits in ascending order.
//
// n will be defined as : 0 < n < 10
func PrintCombN(n int) {
	var combination = []rune{}

	digit := '0'
	generateComb(n, digit, combination)

}
