package main

import (
	"fmt"
	"os"
)

func main() {
	arguments := os.Args[1:]

	if len(arguments) < 1 {
		fmt.Println("File name missing")
	} else if len(arguments) > 1 {
		fmt.Println("Too many arguments")
	} else {

		file, _ := os.Open(string(arguments[0]))

		arr := make([]byte, 15)
		file.Read(arr)
		fmt.Println(string(arr))

		file.Close()

	}
}
