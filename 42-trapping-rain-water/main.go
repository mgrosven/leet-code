package main

import (
	"fmt"
)

func main() {
	a := []int{4, 2, 3}
	r := trap(a)
	fmt.Println(r)
}

func trap(height []int) int {

	total := 0
	if len(height) >= 3 {
		i, j := 0, 1

		trapped := 0
		for i < len(height)-1 && j < len(height) {

			if height[i] > height[j] {
				trapped += height[i] - height[j]
				j++
			} else {
				total += trapped
				trapped = 0
				i = j
				j++
			}
			if j == len(height) {
				i++
				j = i + 1
				trapped = 0
			}
		}
	}
	return total
}
