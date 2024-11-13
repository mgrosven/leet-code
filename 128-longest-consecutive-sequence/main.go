package main

import "fmt"

func main() {
	a := []int{100, 4, 200, 1, 3, 2}
	r := longestConsecutive(a)
	fmt.Printf("Longest Consecutive Sequence: %d", r)
}

func longestConsecutive(nums []int) int {
	m := make(map[int]int)
	longest := 0
	for _, n := range nums {
		m[n] = n
	}

	for n := range m {
		if _, ok := m[n-1]; ok {
			continue
		} else {
			ok = true
			current := 0
			for ok {
				current++
				_, ok = m[n+current]
			}
			if current > longest {
				longest = current
			}
		}
	}
	return longest
}
