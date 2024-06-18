package piscine

// import (
// 	"fmt"
// )

// func hasNumber(s string, r rune) bool {
// 	for _, char := range s {
// 		if char == r {
// 			return true
// 		}
// 	}
// 	return false
// }

// func main() {
// 	p4 := []string{"4th Place"}
// 	p3 := []string{"3rd Place"}
// 	p2 := []string{"2nd Place"}
// 	p1 := []string{"1st Place"}

// 	position := [][]string{p4, p3, p2, p1}
// 	rows := len(position)
// 	cols := len(position[0])
// 	var char rune
// 	// Make an emptry 2d array with same dimensions as podium
// 	newPodium := make([][]string, rows)
// 	for i := range position {
// 		newPodium[i] = make([]string, cols)
// 	}

// 	for range position {
// 		for row, slice := range position {
// 			for col, str := range slice {
// 				char = rune(row+1) + '0'

// 				fmt.Println(str)
// 				fmt.Println(hasNumber(str, char))
// 				if hasNumber(str, char) {
// 					newPodium[row][col] = str
// 				}
// 			}
// 			fmt.Println(char)
// 		}
// 	}
// 	// fmt.Println(len(newPodium))
// 	// fmt.Println(len(newPodium[0]))
// 	fmt.Println(newPodium)

// }
