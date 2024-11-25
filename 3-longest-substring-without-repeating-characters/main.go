package main

import (
	"fmt"
)

func main() {
	s := "abba"
	r := lengthOfLongestSubstring(s)
	fmt.Println(r)
}

func lengthOfLongestSubstring(s string) int {
	longest, left := 0, 0
	m := make(map[rune]int)

	for i, r := range s {
		if v, ok := m[r]; !ok {
			fmt.Println("Doesn't Contain", string(r))
		} else {
			left = max(left, v+1)
			fmt.Println("Resetting left", left)
		}
		m[r] = i
		longest = max(longest, i-left+1)

	}
	return longest
}
