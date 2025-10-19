package main

import (
	"fmt"
)

func main() {
	gain := []int{-5, 1, 5, 0, -7}
	fmt.Println(largestAltitude(gain))
	gain = []int{-4, -3, -2, -1, 4, 3, 2}
	fmt.Println(largestAltitude(gain))
	gain = []int{-1, -2, -3}
	fmt.Println(largestAltitude(gain))
}

func largestAltitude(gain []int) int {
	highest, current := 0, 0

	for _, g := range gain {
		current += g
		if current > highest {
			highest = current
		}
	}
	return highest
}
