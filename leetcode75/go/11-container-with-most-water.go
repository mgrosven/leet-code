package main

import (
	"fmt"
	"math"
)

func maxArea(height []int) int {
	left, right, maxArea := 0, len(height)-1, 0
	for left < right {
		h := math.Min(float64(height[left]), float64(height[right]))
		w := right - left
		area := int(h) * w
		if area > maxArea {
			maxArea = area
		}
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return maxArea
}

func main() {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(maxArea(height))
	height = []int{1, 1}
	fmt.Println(maxArea(height))
}
