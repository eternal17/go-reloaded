package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]

	// error handling

	// if len(args) < 1 {
	// 	fmt.Println("Error, file name is missing")
	// 	os.Exit(0)
	// } else if len(args) == 1 {
	// 	fmt.Println("Error, No file to write to")
	// 	os.Exit(0)
	// } else if len(args) > 2 {
	// 	fmt.Println("Error, too many arguments")
	// 	os.Exit(0)
	// }
	// reading first file

	givenText, _ := os.ReadFile(args[0])

	// array to push the words into
	words := strings.Split(string(givenText), " ")

	// loop over each word in array, satisfy conditions

	for i, word := range words {
		if word == "(up)" {
			words[i-1] = strings.ToUpper(words[i-1])
		}
	}
	fmt.Println(words)

	// man := os.WriteFile(args[1], manipulate, 0644)
	// if man != nil {
	// 	log.Fatal(man)
	// }
}
