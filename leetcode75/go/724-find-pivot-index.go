package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 7, 3, 6, 5, 6}
	result := pivotIndex(nums)
	fmt.Println(result)
	nums = []int{1, 2, 3}
	result = pivotIndex(nums)
	fmt.Println(result)
	nums = []int{2, 1, -1}
	result = pivotIndex(nums)
	fmt.Println(result)
	nums = []int{-1, -1, -1, -1, -1, 0}
	result = pivotIndex(nums)
	fmt.Println(result)
}

func pivotIndex(nums []int) int {
	leftSum, rightSum := 0, 0
	for _, num := range nums {
		rightSum += num
	}

	for i, num := range nums {
		rightSum -= num
		if leftSum == rightSum {
			return i
		}
		leftSum += num
	}
	return -1
}
