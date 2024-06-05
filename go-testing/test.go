package main

import (
	"fmt"
	"piscine"
)

func main() {
	slice := []string{"des  rts", "mains ", "  drinks", "starters  "}
	fmt.Println(piscine.ReverseMenuIndex([]string{"desserts", "mains", "drinks", "starters"}))
	fmt.Println(piscine.ReverseMenuIndex(slice))
	result := []string{"a", "A", "1", "b", "B", "2", "c", "C", "3"}
	piscine.SortWordArr(result)

	fmt.Println(result)
}
