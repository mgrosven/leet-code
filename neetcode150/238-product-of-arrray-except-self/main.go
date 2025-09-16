package main

import "fmt"

func main() {
	a := []int{-1, 1, 0, -3, 0}
	r := productExceptSelf(a)
	fmt.Println(r)
}

func productExceptSelf(nums []int) []int {
	// make array and init all values to 1
	output := make([]int, len(nums))
	for i := range output {
		output[i] = 1
	}
	left, right := 1, 1

	// left pass, i = product of all left values, update total left
	for i, n := range nums {
		output[i] = left
		left *= n
	}

	// right pass, i = product of previous pass and all right values, update total right
	for i := len(nums) - 1; i >= 0; i-- {
		output[i] *= right
		right *= nums[i]
	}
	return output
}
