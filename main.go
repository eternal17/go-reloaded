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
			words = append(words[:i], words[i+1:]...)
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
		} else if word == "(up," {
			b := strings.Trim(string(words[i+1]), words[i+1][1:])
			number, _ := strconv.Atoi(string(b))
			for j := 1; j <= number; j++ {
				words[i-j] = strings.ToUpper(words[i-j])
				words = append(words[:i], words[i+1:]...)
			}
			// lower with number
		} else if word == "(low," {
			b := strings.Trim(string(words[i+1]), words[i+1][1:])
			number, _ := strconv.Atoi(string(b))
			for j := 1; j <= number; j++ {
				words[i-j] = strings.ToLower(words[i-j])
				words = append(words[:i], words[i+1:]...)
			}
			// capitalize with num
		} else if word == "(cap," {
			b := strings.Trim(string(words[i+1]), words[i+1][1:])
			number, _ := strconv.Atoi(string(b))
			for j := 1; j <= number; j++ {
				words[i-j] = strings.Title(words[i-j])
				words = append(words[:i], words[i+1:]...)
			}
			// punctuations
		}
	}
	// RemoveKeywords(words)
	// Punctuations(words)
	ChangeA(words)
	fmt.Println(Punctuations(words))
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
	vowels := []string{"a", "e", "i", "o", "u", "h", "A", "E", "I", "O", "U", "H"}

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

func Punctuations(s []string) []string {
	puncs := []string{",", ".", "!", "?", ":", ";"}
	for i, word := range s {
		for _, punc := range puncs {
			if string(word[0]) == punc && (s[len(s)-1] != s[i]) {
				s[i-1] += punc
				s[i] = word[1:]
			}
			if (string(word[0]) == punc) && (s[len(s)-1] == s[i]) {
				s[i-1] += word
				s = s[:len(s)-1]
			}
		}
	}
	return s
}

// func RemoveKeywords(s []string) []string {
// 	keywords := []string{"(up)", "(low)"}

// 	for _, keyword := range keywords {
// 		for i, word := range s {
// 			if keyword == word {
// 				s = append(s[:i], s[i+1:]...)
// 			}
// 		}
// 	}
// 	return s
// }
