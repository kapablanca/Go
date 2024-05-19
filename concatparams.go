package piscine

// Function that takes the arguments received in parameters
// and returns them as a string. The string is the result of
// all the arguments concatenated with a newline between.
func ConcatParams(args []string) string {
	return Join(args, "\n")
}
