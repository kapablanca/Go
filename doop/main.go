package main

import "os"

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

	valid_operators := [...]string{"+", "-", "/", "*", "%"}

	if len(operator) != 1 {
		return false
	}

	for _, symbol := range valid_operators {
		if operator == symbol {
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
	n := n1 + n2
	return n
}

func subtract(n1, n2 int) int {
	n := n1 - n2
	return n
}

func multiply(n1, n2 int) int {
	n := n1 * n2
	return n
}

func divide(n1, n2 int) int {
	if n2 == 0 {
		return 0
	}
	n := n1 / n2
	return n
}

func modulo(n1, n2 int) int {
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
func applyOperator(f func(int,int)int, n1, n2 int) int {

	result := f(n1, n2)

	return result

}

// Main Program
func main() {

	arguments := os.Args[1:]

	checkArgs(arguments)

	valid_operators := [...]string{"+", "-", "/", "*", "%"}

	list_operations := []func(int, int) int{add, subtract, divide, multiply, modulo}

	n1 := atoi(arguments[0])
	n2 := atoi(arguments[2])

	operation_index := 0

	for index, oper := range valid_operators {
		if arguments[1] == oper {
			operation_index = index
		}
	}

	if operation_index == 4 && n2 == 0 {
		os.Stdout.WriteString("No modulo by 0")
		return 
	}

	if operation_index == 2 && n2 == 0 {
		os.Stdout.WriteString("No modulo by 0")
		return 
	}

	result := applyOperator((list_operations[operation_index], n1, n2))
	os.Stdout.WriteString(result)
}
