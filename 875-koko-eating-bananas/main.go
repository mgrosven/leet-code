package main

import (
	"fmt"
	"math"
)

func main() {
	b := []int{3, 6, 7, 11}
	h := 8
	o := minEatingSpeed(b, h)
	fmt.Println(o)
}

func minEatingSpeed(piles []int, h int) int {
	min, max := 1, getMax(piles)
	k := max

	isPossible := func(k int) bool {
		hours := 0
		for _, pile := range piles {
			hours += int(math.Ceil(float64(pile) / float64(k)))
		}
		return hours <= h
	}

	for min <= max {
		mid := min + (max-min)/2
		if isPossible(mid) {
			k = mid
			max = mid - 1
		} else {
			min = mid + 1
		}
	}
	return k
}

func getMax(nums []int) int {
	max := 0
	for _, n := range nums {
		if n > max {
			max = n
		}
	}
	return max
}
