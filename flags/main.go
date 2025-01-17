package main

import (
	"fmt"
	"os"
)

// Program that can take --insert (or -i), --order (or -o) and
//  a string as arguments. This program should:
// -Insert the string given to the --insert (or -i), into the argument string, if given.
// -If the flag --order (or -o) is given, order the string argument (in ASCII order).
// -If there are no arguments or if the flag --help (or -h) is given, the options should be printed as in the example.
// The short flag will have two spaces before the (-).
// The explanation of the flag will have a tab followed by a space before the beginning of the sentence (This flag...)

// Prints the option of a flag
func printOption(flag, option string) {
	shortFlag := flag[1:3]
	fmt.Println(flag)
	fmt.Println("  " + shortFlag)
	fmt.Println("\t" + " " + option)
}

// Checks for specific flag
func hasFlag(slice []string, flag string) (int, bool) {
	shortFlag := flag[1:3]
	for index, arg := range slice {
		if len(arg) >= len(flag) && arg[:len(flag)] == flag {
			return index, true
		}
		if len(arg) >= len(shortFlag) && arg[:len(shortFlag)] == shortFlag {
			return index, true
		}
	}
	return -1, false
}

// Order a string argument in ASCII order
func orderString(s string) string {
	stringSlice := []rune(s)

	for range stringSlice {
		for i := 0; i < len(stringSlice)-1; i++ {
			if stringSlice[i+1] < stringSlice[i] {
				stringSlice[i], stringSlice[i+1] = stringSlice[i+1], stringSlice[i]
			}
		}
	}
	return string(stringSlice)
}

// Insert a string to a given string
func insertString(flagString, s string) string {
	indexString := 0
	// Finding the index of the string to be inserted
	for index, char := range flagString {
		if char == '=' {
			indexString = index + 1
			break
		}
	}
	return s + flagString[indexString:]
}

// Finding the string argument
func stringArg(args []string, oIndex, iIndex int) string {
	oString := ""
	iString := ""
	answer := ""

	if oIndex >= 0 {
		oString = args[oIndex]
	}
	if iIndex >= 0 {
		iString = args[iIndex]
	}
	// Find the string that is not a flag
	for _, arg := range args {
		if arg != oString && arg != iString {
			answer = arg
		}
	}
	return answer
}

func main() {
	arguments := os.Args[1:]
	help := false
	insert := false
	order := false
	insertIn := -1
	orderIn := -1
	// Check if there are no arguments
	if len(arguments) < 1 {
		help = true
	}

	// Check if there are more arguments
	if len(arguments) > 3 {
		return
	}

	// Declaring the flag variables and options
	flagInsert := "--insert"
	optionI := "This flag inserts the string into the string passed as argument."
	flagOrder := "--order"
	optionO := "This flag will behave like a boolean, if it is called it will order the argument."
	flagHelp := "--help"
	flags := []string{flagInsert, flagOrder}
	options := []string{optionI, optionO}

	// Checking for help flag in arguments
	if !help {
		_, help = hasFlag(arguments, flagHelp)
	}

	// Printing options for each flag
	if help {
		for index, flag := range flags {
			printOption(flag, options[index])
		}
		return
	}

	insertIn, insert = hasFlag(arguments, flagInsert)
	orderIn, order = hasFlag(arguments, flagOrder)
	str := stringArg(arguments, orderIn, insertIn)

	// Insert the argument of flag insert
	if insert {
		str = insertString(arguments[insertIn], str)
	}
	// Order the string
	if order {
		str = orderString(str)
	}

	fmt.Println(str)
}
