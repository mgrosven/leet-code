package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	// s := "A man, a plan, a canal: Panama"
	// s := "race a car"
	s := " "
	r := isPalindrome(s)
	if r {
		fmt.Println("Valid")
	} else {
		fmt.Println("Invalid")
	}
}

func isPalindrome(s string) bool {
	if len(s) < 2 {
		return true
	}
	s = strings.ToLower(s)
	var result []rune
	for _, ch := range s {
		if unicode.IsLetter(ch) || unicode.IsDigit(ch) {
			result = append(result, ch)
		}
	}
	start := 0
	end := len(result) - 1

	for start < end {
		if result[start] != result[end] {
			return false
		}
		start++
		end--
	}

	return true
}
