package main

import "os"

// Functions of operations
func addition(a, b int) (bool, int) {
	valid := true
	result := a + b
	// Overflow check
	if (a > 0 && b > 0) && result < 0 {
		valid = false
	}
	if (a < 0 && b < 0) && result > 0 {
		valid = false
	}
	return valid, result
}

func subtraction(a, b int) (bool, int) {
	valid := true
	result := a - b
	// Overflow check
	if (a < 0 && b > 0) && result > 0 {
		valid = false
	}
	if (a > 0 && b < 0) && result < 0 {
		valid = false
	}
	return valid, result
}

func multiplication(a, b int) (bool, int) {
	valid := true
	result := a * b
	// Overflow check
	if ((a > 0 && b > 0) || (a < 0 && b < 0)) &&
		result < 0 {
		valid = false
	}
	if ((a < 0 && b > 0) || (a > 0 && b < 0)) &&
		result > 0 {
		valid = false
	}

	return valid, result
}

func division(a, b int) (bool, int) {
	valid := true
	result := 0
	// Division by 0 check
	if b == 0 {
		valid = false
		return valid, result
	}
	result = a / b

	return valid, result
}

func modulo(a, b int) (bool, int) {
	valid := true
	result := 0
	// Modulo by 0 check
	if b == 0 {
		valid = false
		return valid, result
	}
	result = a & b
	return valid, result
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

// String to integer
func atoi(s string) (bool, int) {
	sign := 1
	number := 0
	valid := true
	// Find the minus sign
	if s[0] == '-' {
		s = s[1:]
		sign = -1
	}
	// Convert to integer
	for _, char := range s {
		// Check if char is numeric
		if char < '0' || char > '9' {
			valid = false
			break
		}
		char -= '0'
		number = 10*number + int(char)
	}
	// Check if there is an ovreflow
	if (sign > 0 && sign*number < 0) ||
		(sign < 0 && sign*number > 0) {
		valid = false
	}
	return valid, sign * number
}

// Integer to string function
func itoa(number int) string {
	negative := false
	digit := 0
	var char rune
	digits := []rune{}
	// Check the sign
	if number < 0 {
		negative = true
	}
	// Check for number equal to 0
	if number == 0 {
		return "0"
	}
	// Extract the digits
	for number != 0 {
		digit = number % 10
		if negative {
			digit *= -1
		}
		char = rune(digit) + '0'
		digits = append(digits, char)
		number /= 10
	}
	if negative {
		digits = append(digits, '-')
	}
	// Reverse the digits order
	for i, j := 0, len(digits)-1; i < len(digits)/2; i, j = i+1, j-1 {
		digits[i], digits[j] = digits[j], digits[i]
	}
	return string(digits)
}

// Apply function
func apply(s string, a, b int) (bool, int) {
	operations := []func(a, b int) (bool, int){addition, subtraction, multiplication, division, modulo}
	operators := []string{"+", "-", "*", "/", "%"}
	result := 0
	valid := true

	for index, operator := range operators {
		if s == operator {
			valid, result = operations[index](a, b)
			break
		}
	}
	return valid, result
}

// Print function
func print(s string) {
	byteSlice := []byte(s)
	os.Stdout.Write(byteSlice)
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
	validFirst, firstInt := atoi(arguments[0])
	validSecond, secondInt := atoi(arguments[2])
	if !validFirst || !validSecond {
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
	// Apply the appropriate operation
	valid, result := apply(arguments[1], firstInt, secondInt)
	// Check if there was an overflow error
	if !valid {
		return
	}
	// Print the final result
	print(itoa(result))
}
