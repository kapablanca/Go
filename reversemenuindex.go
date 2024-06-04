package piscine

// Function that removes spaces from a string
func trim(s string) string {
	var newString string

	for _, char := range s {
		if char == ' ' {
			continue
		}
		newString += string(char)
	}
	return newString
}

// Function that returns a map with each char of
// the string and the number of times it appears
// Spaces are not counted
func indexChar(s string) map[rune]int {
	m := make(map[rune]int)
	s = trim(s)

	for _, char := range s {
		m[char]++
	}
	return m
}

// Function that checks if two words have most letters the same
// where s1 is the string we compare s2 to see if s2 is similar to s1.
func isSimilar(s1, s2 string) bool {
	map1 := indexChar(s1)
	map2 := indexChar(s2)
	sameChars := 0

	// Check its string map to find the same chars
	for key1 := range map1 {
		for key2, value := range map2 {
			if key2 == key1 {
				sameChars += value
			}
		}
	}
	// Check if s2 has more than half of chars of s1
	return sameChars > len(trim(s1))/2
}

// Function that takes a slice of strings(menu) and returns
// another slice with the items in the correct order.
// append() is not allowed for this exercise.
func ReverseMenuIndex(menu []string) []string {
	menuOrder := []string{"starters", "drinks", "mains", "desserts"}
	correctMenu := make([]string, len(menu))

	for index, item1 := range menuOrder {
		for _, item2 := range menu {
			if trim(item2) == item1 || isSimilar(item1, item2) {
				correctMenu[index] = item2
				break
			}
		}
	}
	return correctMenu
}
