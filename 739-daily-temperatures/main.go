package main

import (
	"fmt"
)

func main() {
	t := []int{30, 60, 90}
	o := dailyTemperatures(t)
	fmt.Println(o)
}

func dailyTemperatures(temperatures []int) []int {
	result := make([]int, len(temperatures))

	tracker := []int{}

	for i, t := range temperatures {
		for len(tracker) > 0 && t > temperatures[tracker[len(tracker)-1]] {
			c := len(tracker) - 1
			result[tracker[c]] = i - tracker[c]
			tracker = tracker[:c]
		}
		tracker = append(tracker, i)
	}
	return result
}
