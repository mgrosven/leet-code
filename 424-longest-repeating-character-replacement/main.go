package main

import "fmt"

func main() {
	s := "AABABBA"
	k := 1
	r := characterReplacement(s, k)
	fmt.Println(r)
}

func characterReplacement(s string, k int) int {
	charCount := make(map[byte]int)
	left, right, maxFreq, maxLen := 0, 0, 0, 0

	for right < len(s) {
		r := s[right]
		charCount[r]++
		maxFreq = max(maxFreq, charCount[r])

		windowLen := right - left + 1
		if windowLen-maxFreq > k {
			charCount[s[left]]--
			left++
		}
		maxLen = max(maxLen, right-left+1)
		right++
	}
	return maxLen
}
