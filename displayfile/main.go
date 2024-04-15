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

		file, _ := os.Open(arguments[0])

		var arr []byte
		file.Read(arr)
		fmt.Println(string(arr))

		file.Close()

	}
}
