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

			// upper with number
		} else if strings.Contains(word, "(up,") {
			b := strings.Trim(string(words[i+1]), ")")
			number, _ := strconv.Atoi(string(b))
			for j := 1; j <= number; j++ {
				words[i-j] = strings.ToUpper(words[i-j])
				words = append(words[:i], words[i+1:]...)
			}
			// lower with number
		} else if strings.Contains(word, "(low,") {
			b := strings.Trim(string(words[i+1]), ")")
			number, _ := strconv.Atoi(string(b))
			for j := 1; j <= number; j++ {
				words[i-j] = strings.ToLower(words[i-j])
				words = append(words[:i], words[i+1:]...)
			}
			// capitalize with num
		} else if strings.Contains(word, "(cap,") {
			b := strings.Trim(string(words[i+1]), ")")
			number, _ := strconv.Atoi(string(b))
			for j := 1; j <= number; j++ {
				words[i-j] = strings.Title(words[i-j])
				words = append(words[:i], words[i+1:]...)
			}
		}
	}
	// Times(words)
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
	vowels := []string{"a", "e", "i", "o", "u", "h", "A", "E", "I", "O", "U"}

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

// func Times(s []string) []string {
// 	keywords := []string{"(up,", "(low,", "(cap,"}

// 	for i, word := range s {
// 		for _, keyword := range keywords {
// 			if word == keyword {
// 				a := strings.Trim(s[i+1], ")")
// 				number, _ := strconv.Atoi(string(a))
// 				for j := 1; j <= number; j++ {
// 					if word == "(up," {
// 						s[i-j] = strings.ToUpper(s[i-j])
// 						s = append(s[:i], s[i+1:]...)
// 					} else if word == "(low," {
// 						s[i-j] = strings.ToLower(s[i-j])
// 					} else if word == "(cap," {
// 						s[i-j] = strings.Title(s[i-j])
// 					}

// 				}
// 			}
// 		}
// 	}
// 	return s
// }
