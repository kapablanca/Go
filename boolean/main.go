package main

import (
	"os"

	"github.com/01-edu/z01"
)

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

func even(nbr int) int {
	if nbr%2 == 0 {
		return 1
	}
	return 0
}

func isEven(nbr int) int {
	if even(nbr) == 1 {
		return 1
	} else {
		return 0
	}
}

func main() {
	EvenMsg := "I have an even number of arguments"
	OddMsg := "I have an odd number of arguments"

	arguments := os.Args[1:]
	lengthOfArg := len(arguments)

	if isEven(lengthOfArg) == 1 {
		printStr(EvenMsg)
	} else {
		printStr(OddMsg)
	}
}