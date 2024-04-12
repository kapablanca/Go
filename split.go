package piscine

func Split(s, sep string) []string {
	var words []string

	for i, word := range s {
	}
}

// var splitted []string
// var word string

// if len(sep) == 1 {
// 	for i := 0; i < len(s); i++ {
// 		if string(s[i]) == sep {
// 			if i != 0 && string(s[i-1]) != sep {

// 				splitted = append(splitted, word)
// 				word = ""
// 			} else {
// 				word = ""
// 			}
// 		} else {

// 			word += string(s[i])

// 			if i == len(s)-1 {
// 				splitted = append(splitted, word)
// 				word = ""
// 			}

// 		}
// 	}
// } else {
// 	for i := 0; i < len(s)-len(sep); i = i + len(sep) {
// 		if string(s[i:i+len(s)]) == sep {
// 			if i != 0 && string(s[i-len(sep):i]) != sep {

// 				splitted = append(splitted, word)
// 				word = ""
// 			} else {
// 				word = ""
// 			}
// 		} else {

// 			word += string(s[i : i+len(s)])

// 			if i == len(s)-len(sep)-1 {
// 				splitted = append(splitted, word)
// 				word = ""
// 			}

// 		}
// 	}
// }

// return splitted
