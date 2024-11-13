package main

import "fmt"

func main() {
	a := []int{-1, 0}
	t := -1
	r := twoSum(a, t)
	fmt.Println(r)
}

func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			total := nums[i] + nums[j]
			if total == target {
				return []int{i + 1, j + 1}
			}
			if total > target {
				break
			}
		}
	}
	return []int{}
}
