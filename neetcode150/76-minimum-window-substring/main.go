package main

import (
	"fmt"
)

func main() {
	s := "a"
	t := "aa"
	r := minWindow(s, t)
	fmt.Println(r)
}

func minWindow(s string, t string) string {

	if len(t) > len(s) {
		return ""
	}

	tcount := make(map[byte]int)
	wcount := make(map[byte]int)

	for i := range t {
		tcount[t[i]]++
	}

	have, need := 0, len(tcount)
	left, minLen, start := 0, len(s)+1, 0

	for right := 0; right < len(s); right++ {
		char := s[right]
		wcount[char]++

		if tcount[char] > 0 && wcount[char] == tcount[char] {
			have++
		}

		for have == need {
			if right-left+1 < minLen {
				minLen = right - left + 1
				start = left
			}

			wcount[s[left]]--
			if tcount[s[left]] > 0 && wcount[s[left]] < tcount[s[left]] {
				have--
			}
			left++
		}
	}

	if minLen > len(s) {
		return ""
	}
	return s[start : start+minLen]
}
