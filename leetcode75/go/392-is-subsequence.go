package main

import "fmt"

func main() {
	s := "abc"
	t := "ahbgdc"
	fmt.Println(isSubsequence(s, t))
	s = "axc"
	t = "ahbgdc"
	fmt.Println(isSubsequence(s, t))
}

func isSubsequence(s string, t string) bool {
	sLen, tLen := len(s), len(t)
	if sLen == 0 {
		return true
	}
	sIndex, tIndex := 0, 0
	for tIndex < tLen && sIndex < sLen {
		if s[sIndex] == t[tIndex] {
			sIndex++
		}
		tIndex++
	}
	return sIndex == sLen
}
