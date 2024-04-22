package main

import (
	"os"
)

// Maximum value of integers
const (
	MaxInt = 9223372036854775807
	MinInt = -9223372036854775808
)

// Checking is string is numerical value
func isNumber(s string) bool {
	for _, char := range s {
		if char < '0' || char > '9' {
			return false
		}
	}
	return true
}

// Checking validity of operator
func checkOperator(operator string) bool {
	// valid_operators := [...]string{"+", "-", "/", "*", "%"}

	valid_operators := [...]rune{'+', '-', '/', '*', '%'}

	if len(operator) != 1 {
		return false
	}

	for _, symbol := range valid_operators {
		if operator == string(symbol) {
			return true
		}
	}
	return false
}

// Sanitizing arguments
func checkArgs(args []string) bool {
	if len(args) != 3 {
		return false
	}
	if !isNumber(args[0]) || !isNumber(args[2]) {
		return false
	}
	if !checkOperator(args[1]) {
		return false
	}

	return true
}

// Mathematical operations functions
func add(n1, n2 int) int {
	if n1 > MaxInt-n2 || n1 < MinInt-n2 {
		os.Exit(0)
	}
	n := n1 + n2
	return n
}

func subtract(n1, n2 int) int {
	if n1 > MaxInt+n2 || n1 < MinInt+n2 {
		os.Exit(0)
	}
	n := n1 - n2
	return n
}

func multiply(n1, n2 int) int {
	if n1 > MaxInt/n2 || n1 < MinInt/n2 {
		os.Exit(0)
	}
	n := n1 * n2
	return n
}

func divide(n1, n2 int) int {
	if n2 == 0 {
		os.Stdout.WriteString("No division by 0\n")
		os.Exit(0)
	}
	n := n1 / n2
	return n
}

func modulo(n1, n2 int) int {
	if n2 == 0 {
		os.Stdout.WriteString("No modulo by 0\n")
		os.Exit(0)
	}
	n := n1 % n2
	return n
}

// Converting string to a number
func atoi(s string) int {
	number := 0
	sign := 1
	start := 0

	if s[0] == '-' || s[0] == '+' {
		if s[0] == '-' {
			sign = -1
		}
		start = 1
	}

	for i := start; i < len(s); i++ {
		digit := s[i] - '0'
		number = number*10 + int(digit)
	}
	return number * sign
}

// Function to apply the selected operator
func applyOperator(f func(int, int) int, n1, n2 int) int {
	result := f(n1, n2)

	return result
}

// Convert number to string
func itoa(num int) string {
	// Special case for zero
	if num == 0 {
		return "0"
	}

	// Check if the number is negative
	isNegative := num < 0
	if isNegative {
		num = -num // Make the number positive for processing
	}

	// This will store the digits of the number.
	digits := []byte{}

	// Extract digits of the number
	for num > 0 {
		digit := num % 10
		// Prepend digit to the slice (since we're reading the digits in reverse order)
		digits = append([]byte{byte('0' + digit)}, digits...)
		num /= 10
	}

	// If the number was negative, prepend a '-' sign
	if isNegative {
		digits = append([]byte{'-'}, digits...)
	}

	// Convert byte slice to string
	return string(digits)
}

// Main Program
func main() {
	if len(os.Args) < 2 {
		return
	}
	arguments := os.Args[1:]

	if !checkArgs(arguments) {
		return
	}

	valid_operators := [...]rune{'+', '-', '/', '*', '%'}

	list_operations := []func(int, int) int{add, subtract, divide, multiply, modulo}

	n1 := atoi(arguments[0])
	n2 := atoi(arguments[2])

	operation_index := 0

	for index, oper := range valid_operators {
		if arguments[1] == string(oper) {
			operation_index = index
		}
	}

	if operation_index == 4 && n2 == 0 {
		os.Stdout.WriteString("No modulo by 0")
		return
	}

	if operation_index == 2 && n2 == 0 {
		os.Stdout.WriteString("No division by 0")
		return
	}

	result := applyOperator(list_operations[operation_index], n1, n2)
	os.Stdout.WriteString(itoa(result) + "\n")
}
