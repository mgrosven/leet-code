package main

import "fmt"

func main() {
	a := []int{1,2,3,1}
	r := containsDuplicate(a)
	if (r) {
		fmt.Println("Has Duplicate")
	} else {
		fmt.Println("Does Not Have Duplicate")
	}
}

func containsDuplicate(nums []int) bool {
    m := make(map[int]bool)
	hasDuplicate := false
	for _, n := range nums {
		v := m[n]
		if (v) {
			hasDuplicate = true
			break
		}
		m[n] = true
	}
	return hasDuplicate
}
