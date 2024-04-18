package main

import (
	"fmt"
	"os"
)

func main() {
	arguments := os.Args[1:]

	for _, arg := range arguments {
		if arg == "01" || arg == "galaxy" || arg == "galaxy 01" {
			fmt.Println("Alert!!!")
		}
	}
}
