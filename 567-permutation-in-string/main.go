package main

import "fmt"

func main() {
	s1 := "ab"
	s2 := "eidboaoo"
	r := checkInclusion(s1, s2)
	if r {
		fmt.Println("Contains Permuation")
	} else {
		fmt.Println("Does Not Contain Permuation")
	}
}

func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}
	countS1 := make([]int, 26)
	windowCount := make([]int, 26)

	for _, r := range s1 {
		countS1[r-'a']++
	}

	for i := 0; i < len(s1); i++ {
		windowCount[s2[i]-'a']++
	}

	if matches(countS1, windowCount) {
		return true
	}

	for i := len(s1); i < len(s2); i++ {
		// Add the new character to the window
		windowCount[s2[i]-'a']++
		// Remove the first character of the previous window
		windowCount[s2[i-len(s1)]-'a']--
		// Check if the current window matches
		if matches(countS1, windowCount) {
			return true
		}
	}

	return false
}

func matches(a, b []int) bool {
	for i := 0; i < 26; i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
