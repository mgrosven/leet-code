package main

import (
	"fmt"
)

func main() {

	l1 := []int{3, 3, 3, 3, 3}

	o := findDuplicate(l1)
	fmt.Println(o)
}

func findDuplicate(nums []int) int {
	m := make(map[int]int)

	for _, n := range nums {
		if _, ok := m[n]; ok {
			return n
		} else {
			m[n] = 1
		}
	}
	return -1
}
