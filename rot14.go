package piscine

// Function that transforms a string. Each letter wiil be replaced
// by the letter 14 spots ahead in the alphabetical order. The case
// of the letters stays the same
func Rot14(s string) string {
	rotatedSlice := []rune{}
	var rotatedChar rune
	// Variables used in our calculations
	letterRange := rune(26)
	rot := rune(14)
	lowerLetterMod := 'a' - rot
	upperLetterMod := 'A' - rot

	for _, char := range s {
		if !isLetter(char) {
			rotatedChar = char
			//  (char - 'a' + rot) initiates the letters from 0-25. Adding
			//  the rot value and taking the mod of it allows to cycle through the
			// letters if the range goes beyond 26. Adding the 'a' lastly gives us the
			// desired rune
		} else if IsLower(string(char)) {
			rotatedChar = (char-lowerLetterMod)%letterRange + 'a'
			// Same procedure with lower letters, only changing 'a' to 'A' for uppercase
		} else {
			rotatedChar = (char-upperLetterMod)%letterRange + 'A'
		}
		rotatedSlice = append(rotatedSlice, rotatedChar)
	}
	return string(rotatedSlice)
}
