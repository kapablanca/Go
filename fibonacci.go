package piscine

/*
Function that recursively calculates the value

at the position index in the fibonacci sequence.

A negative index will return -1
*/
func Fibonacci(index int) int {
	// Checking if index is negative
	if index < 0 {
		return -1
	}

	if index == 0 {
		return 0
	}
	if index == 1 {
		return 1
	}
	return Fibonacci(index-1) + Fibonacci(index-2)
}
