package main

import "fmt"

func main() {
	r := isAnagram("anagram", "nagaram")
	if r {
		fmt.Println("Is Anagram")
	} else {
		fmt.Println("Is Not Anagram")
	}
}

func isAnagram(s string, t string) bool {
	m := make(map[rune]int)

	if len(s) != len(t) {
		return false
	}

	for _, r := range s {
		m[r]++
	}

	for _, r := range t {
		m[r]--
	}

	for _, v := range m {
		if v != 0 {
			return false
		}
	}

	return true
}
