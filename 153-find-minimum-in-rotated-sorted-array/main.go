package main

import (
	"fmt"
)

func main() {
	a := []int{2, 1}
	o := findMin(a)
	fmt.Println(o)
}

func findMin(nums []int) int {
	left, right := 0, len(nums)-1

	for left < right {
		mid := left + (right-left)/2
		if nums[mid] > nums[right] {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return nums[left]
}
