package main

import (
	"os"

	"github.com/01-edu/z01"
)

// Print function
func print(s string) {
	for _, char := range s {
		z01.PrintRune(char)
	}
}

// Program that behaves like a simplified cat command. If the programm
//
//	is called without arguments, it should take the standard input(stding) and
//
// print it back on the standard output(stdout).
func main() {
	arguments := os.Args[1:]
	// If no arguments, print back the Stdin
	if len(arguments) < 1 {
		arr := make([]byte, 1)
		for {
			_, err := os.Stdin.Read(arr)
			if err != nil {
				break
			}
			print(string(arr))
		}
	}

	for _, arg := range arguments {
		content, err := os.ReadFile(arg)
		// Print the error
		if err != nil {
			print("ERROR: " + err.Error())
			continue
		}
		print(string(content))
	}
}
