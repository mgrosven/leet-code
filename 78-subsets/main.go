package main

import "fmt"

func main() {
	s := []int{1, 2, 3}

	fmt.Println(subsets(s))
}

func subsets(nums []int) [][]int {
	var res [][]int
	backtrack(&res, []int{}, nums, 0)
	return res
}

func backtrack(result *[][]int, current []int, nums []int, start int) {
	subsetCopy := make([]int, len(current))
	copy(subsetCopy, current)
	*result = append(*result, subsetCopy)

	for i := start; i < len(nums); i++ {
		current = append(current, nums[i])
		backtrack(result, current, nums, i+1)
		current = current[:len(current)-1]
	}
}
