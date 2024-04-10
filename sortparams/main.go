package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args[1:]
	table := arguments

	for i := 0; i < len(arguments)-1; i++ {
		for j := 1 + i; j < len(arguments); j++ {
			if table[i][0] > table[j][0] {
				table[i], table[j] = table[j], table[i]
			}
		}
	}

	for _, character := range table {
		z01.PrintRune(rune(character[0]))
		z01.PrintRune('\n')
	}
}
