package main

import "fmt"

func main() {
	word1 := "abc"
	word2 := "pqr"
	result := mergeAlternately(word1, word2)
	fmt.Println(result)
}

func mergeAlternately(word1 string, word2 string) string {
	merged := ""
	for len(word1) > 0 || len(word2) > 0 {
		if len(word1) > 0 {
			merged += string(word1[0])
			word1 = word1[1:]
		}
		if len(word2) > 0 {
			merged += string(word2[0])
			word2 = word2[1:]
		}
	}
	return merged
}
