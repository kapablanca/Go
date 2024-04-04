package piscine

import "fmt"

// PrintCombN prints all possible combinations of n different digits in ascending order
func PrintCombN(n int) {
	if n < 1 {
		return
	}
	printCombinationRecursively(n, "", 0)
	fmt.Println() // Print a newline character at the end
}

// Recursive function to generate combinations
func printCombinationRecursively(n int, prefix string, start int) {
	if n == 0 {
		fmt.Print(prefix)
		if len(prefix) < 10-n {
			fmt.Print(", ")
		}
		return
	}

	for i := start; i <= 9; i++ {
		newPrefix := prefix + fmt.Sprintf("%d", i)
		printCombinationRecursively(n-1, newPrefix, i+1)
	}
}
