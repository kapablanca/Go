package main

import "os"

// Functions of operations
func addition(a, b int) int {
	return a + b
}

func subtraction(a, b int) int {
	return a - b
}

func multiplication(a, b int) int {
	return a * b
}

func division(a, b int) int {
	return a / b
}

func modulo(a, b int) int {
	return a % b
}

// ////////////////////////////////////////////////
// Check the operator
func validOperator(s string) bool {
	operators := []string{"+", "-", "*", "/", "%"}

	for _, operator := range operators {
		if s == operator {
			return true
		}
	}
	return false
}

// Check division or modulo

// Apply function
func apply(s string) int {
	operations := []func(a, b int) int{addition, subtraction, multiplication, division, modulo}
	operators := []string{"+", "-", "*", "/", "%"}

	for index, operator := range operators {
		if s == operator {

		}
	}

}

// String to integer
func atoi(s string) int {
	if s[0] == '-' {
		s = s[1:]
	}
	number := 0
	for _, char := range s {
		char -= '0'
		number = 10*number + int(char)
	}
	return number
}

// Valid integer
func isInteger(s string) bool {
	for index, char := range s {
		if index == 1 && char == '-' {
			continue
		}
		if char < '0' || char > '9' {
			return false
		}
	}
	return true
}

// Program that has to be used with three arguments:
// a value, an operator(+, -, *, /, %) and another value
// In case of invalid operator, value of arguments or overflow,
// program prints nothing. The program has to handle the modulo and division
// operations by 0 with a message.
func main() {
	arguments := os.Args[1:]
	// Check number of arguments
	if len(arguments) != 3 {
		return
	}
	// Check valid operator
	if !validOperator(arguments[1]) {
		return
	}
	// Check valid integers
	if !isInteger(arguments[0]) || !isInteger(arguments[2]) {
		return
	}
	// Check division and modulo by 0
	if arguments[2] == "0" {
		if arguments[1] == "/" {
			byteDiv := []byte("No division by 0")
			os.Stdout.Write(byteDiv)
			return
		}
		if arguments[1] == "%" {
			byteMod := []byte("No modulo by 0")
			os.Stdout.Write(byteMod)
			return
		}
	}

}
