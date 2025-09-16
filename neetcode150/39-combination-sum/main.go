package main

import "fmt"

func main() {
	c := []int{2, 3, 6, 7}
	t := 7
	fmt.Println(combinationSum(c, t))
}

func combinationSum(candidates []int, target int) [][]int {
	var res [][]int
	backtrack(&res, []int{}, candidates, 0, target)
	return res
}

func backtrack(result *[][]int, current []int, candidates []int, index, target int) {
	if 0 == target {
		subsetCopy := make([]int, len(current))
		copy(subsetCopy, current)
		*result = append(*result, subsetCopy)
		return
	}

	for i := index; i < len(candidates); i++ {
		if candidates[i] <= target {
			current = append(current, candidates[i])
			backtrack(result, current, candidates, i, target-candidates[i])
			current = current[:len(current)-1]
		}
	}
}
