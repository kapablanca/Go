package piscine

// Function that takes a string and returns another string
// with every third character.
// Prints the output followed by a newline.
// If the string is empty, return newline.
// If there is no third character, return newline.
func JumpOver(str string) string {
	newString := []rune{}
	// Check if string is valid
	if len(str) < 3 {
		return "\n"
	}
	// Extract every third char from string
	for index, char := range str {
		if (index+1)%3 == 0 {
			newString = append(newString, char)
		}
	}
	// Add newline at the end
	newString = append(newString, '\n')
	return string(newString)
}
