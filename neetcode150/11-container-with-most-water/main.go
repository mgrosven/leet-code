package main

import (
	"fmt"
)

func main() {
	a := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	r := maxArea(a)
	fmt.Println(r)
}

func maxArea(height []int) int {
	maxArea := 0
	for left, right := 0, len(height)-1; left < right; {
		area := (right - left) * min(height[left], height[right])
		maxArea = max(maxArea, area)
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return maxArea
}
