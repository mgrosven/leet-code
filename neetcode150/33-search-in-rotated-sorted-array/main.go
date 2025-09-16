package main

import (
	"fmt"
)

func main() {
	nums := []int{3, 1}
	t := 1
	o := search(nums, t)
	fmt.Println(o)
}

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := (right + left) / 2
		if nums[mid] == target {
			return mid
		} else if nums[left] <= nums[mid] {
			if target > nums[mid] || target < nums[left] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		} else {
			if target < nums[mid] || target > nums[right] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
	}

	return -1
}
