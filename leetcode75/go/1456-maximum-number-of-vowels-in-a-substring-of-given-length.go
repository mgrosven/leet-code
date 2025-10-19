package main

import (
	"fmt"
)

func main() {
	s := "abciiidef"
	k := 3
	result := maxVowels(s, k)
	fmt.Println(result) // Output: 3
}

func maxVowels(s string, k int) int {
	maxVowels := 0
	currentVowels := 0
	var vowels [128]bool
	vowels['a'], vowels['e'], vowels['i'], vowels['o'], vowels['u'] = true, true, true, true, true

	for i := 0; i < k; i++ {
		if vowels[s[i]] {
			currentVowels++
		}
	}

	maxVowels = currentVowels

	for i := k; i < len(s); i++ {
		if vowels[s[i]] {
			currentVowels++
		}
		if vowels[s[i-k]] {
			currentVowels--
		}
		if currentVowels > maxVowels {
			maxVowels = currentVowels
		}
	}

	return maxVowels
}
