package main

import (
	"fmt"
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
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		for i < j && !(unicode.IsLetter(rune(s[i])) || unicode.IsDigit(rune(s[i]))) {
			i++
		}
		for i < j && !(unicode.IsLetter(rune(s[j])) || unicode.IsDigit(rune(s[j]))) {
			j--
		}
		if i == j {
			break
		}

		if unicode.ToLower(rune(s[i])) != unicode.ToLower(rune(s[j])) {
			return false
		}
	}
	return true
}
