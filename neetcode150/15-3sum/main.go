package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []int{0, 0, 0}
	r := threeSum(a)
	fmt.Println(r)
}

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	result := [][]int{}

	for i := range nums {
		if nums[i] >= 1 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left, right := i+1, len(nums)-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum == 0 {
				result = append(result, []int{nums[i], nums[left], nums[right]})
				for left++; left < right && nums[left] == nums[left-1]; {
					left++
				}
				for right--; left < right && nums[right] == nums[right+1]; {
					right--
				}
			} else if sum < 0 {
				left++
			} else {
				right--
			}
		}
	}
	return result
}
