package main

import (
	"fmt"
)

func main() {
	s := "-15"
	if s < "1" || s > "26" {
		fmt.Println("It is NOT valid")
	} else {
		fmt.Println("It is valid")
	}
}
