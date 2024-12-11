package main

import (
	"fmt"
)

func main() {
	h := []int{2, 1, 5, 6, 2, 3}
	o := largestRectangleArea(h)
	fmt.Println(o)
}

func largestRectangleArea(heights []int) int {
	stack := []int{} // Initialize stack with a dummy index
	maxArea := 0
	n := len(heights)

	for i := 0; i <= n; i++ {
		ch := 0
		if i < n {
			ch = heights[i]
		}

		for len(stack) > 0 && heights[stack[len(stack)-1]] > ch {
			height := heights[stack[len(stack)-1]]
			stack = stack[:len(stack)-1]
			width := i
			if len(stack) > 0 {
				width = i - stack[len(stack)-1] - 1
			}
			maxArea = max(maxArea, height*width)
		}
		stack = append(stack, i)
	}

	return maxArea
}
