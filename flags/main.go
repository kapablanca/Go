package main

import (
	"fmt"
	"os"
)

func main() {
	arguments := os.Args[1:]

	// Stating the option strings
	insert_string := "--insert\n" + "  -i"
	order_string := "--order\n" + "  -o"
	insert_description := "This flag inserts the string into the string passed as argument."
	order_description := "This flag will behave like a boolean, if it is called it will order the argument."

	// Stating the option keywords
	// long_insert := "--insert"
	// short_insert := "-i"
	long_order := "--order"
	short_order := "-o"
	// Stating the variables in use
	is_order := false
	is_insert := false
	index_insert := 0
	index_string := 0
	var answer string

	// Checking if the help option was selected
	if len(arguments) == 0 || arguments[0] == "--help" || arguments[0] == "-h" {

		fmt.Println(insert_string)
		fmt.Println("\t " + insert_description)
		fmt.Println(order_string)
		fmt.Println("\t " + order_description)

	} else {
		// Print the string argument
		if len(arguments) == 1 {

			answer = arguments[0]
			fmt.Println(answer)

		} else {
			// Check for the flags and where they are
			for i, arg := range arguments {

				if IsInsert(arg) {
					is_insert = true
					index_insert = i
				}

				if arg == long_order || arg == short_order {
					is_order = true
				}
				// Saving the place of the string argument
				index_string = i
			}
		}
		// Checking insert instruction
		if is_insert {
			answer = InsertArg(string(arguments[index_insert]), string(arguments[index_string]))
		} else {
			// Inserting just the sting argument to return
			answer = arguments[index_string]
		}
		// Checking order instruction
		if is_order {
			answer = OrderArg(answer)
		}
		// Printing the final answer
		fmt.Println(answer)
	}
}

// Inserting function
func InsertArg(arg string, extra string) string {
	var answer string
	index := 0
	// Checking the index of where the equal sign is
	for i, char := range arg {
		if char == '=' {
			index = i
		}
	}
	// Extracting the string argument to be inserted
	arg = arg[index+1:]

	answer = extra + arg

	return answer
}

// Ordering function
func OrderArg(arg string) string {
	var answer string
	arguments := []rune(arg)
	// Ordering algorithm
	for i := 0; i < len(arguments); i++ {
		for j := 0; j < len(arguments)-1; j++ {
			if arguments[j] > arguments[j+1] {
				arguments[j], arguments[j+1] = arguments[j+1], arguments[j]
			}
		}
	}
	// Converting []rune to string
	for _, char := range arguments {
		answer += string(char)
	}

	return answer
}

// Checking if there is insert inside an argument
func IsInsert(s string) bool {
	for _, char := range s {
		if char == '=' {
			return true
		}
	}
	return false
}
