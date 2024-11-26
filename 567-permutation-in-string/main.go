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
	l, count := 0, [26]int{}

	// set up counts for s1
	for i := range s1 {
		count[s1[i]-'a']++
	}

	// build and check window
	for r := range s2 {
		// reduce current letter
		count[s2[r]-'a']--

		// if all 0, permutation found
		if count == [26]int{} {
			return true
		}

		// if window too big, re-add first character and shift window
		if r+1 > len(s1) {
			count[s2[l]-'a']++
			l++
		}
	}
	return false
}
