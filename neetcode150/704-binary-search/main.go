package main

import (
	"fmt"
)

func main() {
	n := []int{-1, 0, 3, 5, 9, 12}
	t := 9
	o := search(n, t)
	fmt.Println(o)
}

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		v := nums[mid]
		if v == target {
			return mid
		} else if v < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}
