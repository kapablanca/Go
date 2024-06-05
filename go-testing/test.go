package main

import (
	"fmt"
	"piscine"
)

func main() {
	slice := []string{"des  rts", "mains ", "  drinks", "starters  "}
	fmt.Println(piscine.ReverseMenuIndex([]string{"desserts", "mains", "drinks", "starters"}))
	fmt.Println(piscine.ReverseMenuIndex(slice))
}
