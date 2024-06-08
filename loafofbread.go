package piscine

// Function that removes spaces from the string. If a space is in the place
// where the character would be skipped from LoadOfBread(), the returned string
// has a '6' in its place.
func trim(s string) string {
	countSpaces := 0
	countChars := 0
	newString := []rune{}
	// Check each char of the string
	for _, char := range s {
		countChars += 1
		// If the space would be skipped
		if char == ' ' && (countChars-countSpaces)%6 == 0 {
			newString = append(newString, '6')
			countSpaces += 1
			// If the space is in the middle of the string
		} else if char == ' ' {
			countSpaces += 1
			continue
			// Append any other char
		} else {
			newString = append(newString, char)
		}
	}
	return string(newString)
}

// Function that takes a string and returns another one with words of
// 5 characters and skips the next character followed by newline.
// If there is space in the middle of a string it should ignore it and
// get the next character until getting to a length of 5.
// If the string is less thatn 5 characters return "Invalid Output\n"
func LoafOfBread(str string) string {
	// Check if string has less than 5 characters
	if len(str) < 5 {
		return "Invalid Output\n"
	}
	// Remove the middle spaces from string
	str = trim(str)
	// Declare variables
	count := 0
	word := []rune{}
	sentence := []rune{}
	//Loop through the string to create new 5 char words
	for _, char := range str {
		count += 1
		// If we reached the 5 chars to format a word
		if count > 5 {
			count = 0
			word = append(word, ' ')
			sentence = append(sentence, word...)
			word = word[:0]
			continue
		} else {
			word = append(word, char)
		}
	}
	// Append the last word to the final string
	if len(word) != 0 {
		sentence = append(sentence, word...)
	}
	return string(sentence) + "\n"
}
