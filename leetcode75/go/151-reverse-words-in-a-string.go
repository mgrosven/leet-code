package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "  hello  world  "
	result := reverseWords(s)
	fmt.Println(result)
}

func reverseWords(s string) string {
	s = strings.TrimSpace(s)
	words := strings.Fields(s)
	left, right := 0, len(words)-1
	for left < right {
		words[left], words[right] = words[right], words[left]
		left++
		right--
	}
	return strings.Join(words, " ")
}
