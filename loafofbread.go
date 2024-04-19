package piscine

func LoafOfBread(str string) string {
	if str == "" {
		return "\n"
	}
	if countNonSpaceChars(str) < 5 {
		return "Invalid Output\n"
	}

	var result string
	nonSpaceCount := 0 // Counts non-space characters up to 5
	skipNext := false  // Flag to skip the next character

	for i, char := range str {
		if skipNext {
			skipNext = false // Reset the skip flag after skipping one character
			continue         // Skip this character
		}

		if char == ' ' {
			continue // Skip spaces
		}

		nonSpaceCount++
		if nonSpaceCount <= 5 {
			result += string(char)
		}

		if nonSpaceCount == 5 {
			if i+1 < len(str) { // Ensure there is a character to skip
				skipNext = true // Set skip flag to skip the next character
				result += " "   // Add a space to separate the groups
			}
			nonSpaceCount = 0 // Reset the non-space character counter
		}
	}

	// Trim the last space if added and return the result with a newline
	if len(result) > 0 && result[len(result)-1] == ' ' {
		result = result[:len(result)-1]
	}
	return result + "\n"
}

// Helper function to count non-space characters
func countNonSpaceChars(s string) int {
	count := 0
	for _, char := range s {
		if char != ' ' {
			count++
		}
	}
	return count
}
