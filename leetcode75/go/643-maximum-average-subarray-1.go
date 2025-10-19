package main

import (
	"fmt"
	"math"
)

func main() {
	nums := []int{1, 12, -5, -6, 50, 3}
	k := 4
	fmt.Println(findMaxAverage(nums, k))

	nums2 := []int{5}
	k2 := 1
	fmt.Println(findMaxAverage(nums2, k2))
}

func findMaxAverage(nums []int, k int) float64 {
	maxAverage := float64(0)
	windowSum := 0

	for i := 0; i < k; i++ {
		windowSum += nums[i]
	}

	maxAverage = float64(windowSum) / float64(k)

	for i := k; i < len(nums); i++ {
		windowSum += nums[i] - nums[i-k]
		currentAverage := float64(windowSum) / float64(k)
		maxAverage = math.Max(maxAverage, currentAverage)
	}

	return maxAverage
}
