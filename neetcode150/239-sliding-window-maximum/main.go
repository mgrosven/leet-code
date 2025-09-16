package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	k := 3
	r := maxSlidingWindow(nums, k)
	fmt.Println(r)
}

func maxSlidingWindow(nums []int, k int) []int {
	deque := []int{}
	o := []int{}

	for right := 0; right < len(nums); right++ {

		if len(deque) > 0 && deque[0] < right-k+1 {
			deque = deque[1:]
		}

		for len(deque) > 0 && nums[deque[len(deque)-1]] < nums[right] {
			deque = deque[:len(deque)-1]
		}

		deque = append(deque, right)

		if right >= k-1 {
			o = append(o, nums[deque[0]])
		}
	}
	return o
}
