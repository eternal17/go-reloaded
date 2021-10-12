package main

import (
	"fmt"
	"os"
	"strconv"
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
		// fl := string(words[i+1][0])
		if word == "(up)" {
			words[i-1] = strings.ToUpper(words[i-1])
			words = append(words[:i], words[i+1:]...)
		} else if word == "(low)" {
			words[i-1] = strings.ToLower(words[i-1])
		} else if word == "(cap)" {
			words[i-1] = strings.Title(words[i-1])
			words = append(words[:i], words[i+1:]...)
		} else if word == "(hex)" {
			words[i-1] = HextoInt(words[i-1])
			words = append(words[:i], words[i+1:]...)
		} else if word == "(bin)" {
			words[i-1] = BintoInt(words[i-1])
			words = append(words[:i], words[i+1:]...)
		} else if word == "a" && (string(words[i+1][0]) == "a") {
			words[i] = "an"
		}
	}

	// removeElement(words)
	fmt.Println(words)

	// man := os.WriteFile(args[1], manipulate, 0644)
	// if man != nil {
	// 	log.Fatal(man)
	// }
}

// conv hex to int
func HextoInt(hex string) string {
	number, _ := strconv.ParseInt(hex, 16, 64)
	return fmt.Sprint(number)
}

// conv binary to int

func BintoInt(bin string) string {
	number, _ := strconv.ParseInt(bin, 2, 64)
	return fmt.Sprint(number)
}

// func checkVowel(a string) bool {
// 	vowels := []string{"a", "e", "i", "o", "u", "h"}

// 	for _, ch := range vowels {
// 		if a == ch {
// 			return true
// 		}
// 	}
// 	return false
// }

// func removeElement(a []string) []string {
// 	keyWords := []string{"(bin)", "(hex)", "(up)"}
// 	for i, word := range a {
// 		for _, keyword := range keyWords {
// 			if word == keyword {
// 				a = append(a[:i], a[i+1:]...)
// 			}
// 		}
// 	}
// 	return a
// }

// fl == "a" || fl == "e" || fl == "i" || fl == "o" || fl == "u" || fl == "h"
