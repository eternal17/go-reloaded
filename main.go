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
		}
	}

	ChangeA(words)
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

func ChangeA(s []string) []string {
	vowels := []string{"a", "e", "i", "o", "u", "h", ""}

	for i, word := range s {
		for _, letter := range vowels {
			if word == "a" && string(s[i+1][0]) == letter {
				s[i] = "an"
			} else if word == "A" && string(s[i+1][0]) == letter {
				s[i] = "An"
			}
		}
	}
	return s
}

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
