package main

import (
	"fmt"
	"os"
)

// Program that displays on the standard output  Alert!!!
// followed by newline if at least one of the arguments passed in parameter
//
//	matches the string: 01, galaxy or galaxy 01. If none of the parameters
//
// match, the program displays nothing
func main() {
	arguments := os.Args[1:]
	// Check if there are enough arguments
	if len(arguments) < 1 {
		return
	}
	// Check the arguments for the keywords : "01", "galaxy", "galaxy 01"
	for _, arg := range arguments {
		if arg == "01" || arg == "galaxy" || arg == "galaxy 01" {
			fmt.Println("Alert!!!")
			break
		}
	}
}
