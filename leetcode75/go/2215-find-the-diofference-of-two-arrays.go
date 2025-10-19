package main

import (
	"fmt"
)

func main() {
	nums1 := []int{1, 2, 3}
	nums2 := []int{2, 4, 6}
	result := findDifference(nums1, nums2)
	fmt.Println(result)

	nums1 = []int{1, 2, 3, 3}
	nums2 = []int{1, 1, 2, 2}
	result = findDifference(nums1, nums2)
	fmt.Println(result)
}

func findDifference(nums1 []int, nums2 []int) [][]int {
	s1 := make(map[int]struct{})
	s2 := make(map[int]struct{})

	for _, num := range nums1 {
		s1[num] = struct{}{}
	}
	for _, num := range nums2 {
		s2[num] = struct{}{}
	}

	diff1 := []int{}
	diff2 := []int{}

	for num := range s1 {
		if _, found := s2[num]; !found {
			diff1 = append(diff1, num)
		}
	}
	for num := range s2 {
		if _, found := s1[num]; !found {
			diff2 = append(diff2, num)
		}
	}

	return [][]int{diff1, diff2}
}
