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

	total := 0
	l := len(height)
	if l >= 3 {
		maxLeft := make([]int, l)
		maxRight := make([]int, l)

		maxLeft[0] = height[0]
		for i := 1; i < l; i++ {
			maxLeft[i] = max(maxLeft[i-1], height[i])
		}

		maxRight[l-1] = height[l-1]
		for j := l - 2; j >= 0; j-- {
			maxRight[j] = max(maxRight[j+1], height[j])
		}

		for i := 0; i < l; i++ {
			total += max(0, min(maxLeft[i], maxRight[i])-height[i])
		}
	}
	return total
}
