package main

import (
	"fmt"
)

func main() {
	a := []int{4, 2, 0, 3, 2, 5}
	r := trap(a)
	fmt.Println(r)
}

func trap(height []int) int {

	left, right := 0, len(height)-1
	leftMax, rightMax, water := height[left], height[right], 0
	for left < right {
		if leftMax < rightMax {
			left++
			if height[left] > leftMax {
				leftMax = height[left]
			}
			water += leftMax - height[left]
		} else {
			right--
			if height[right] > rightMax {
				rightMax = height[right]
			}
			water += rightMax - height[right]
		}
	}
	return water
}
