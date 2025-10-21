package main

import (
	"fmt"
)

func main() {
	n := 10
	result := guessNumber(n)
	fmt.Println(result) // Output will depend on the implementation of guess API
}

func guessNumber(n int) int {
	left, right := 1, n
	for left <= right {
		mid := left + (right-left)/2
		guessResult := guess(mid)
		if guessResult == 0 {
			return mid
		} else if guessResult < 0 {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}

func guess(num int) int {
	// This is a placeholder for the actual guess API.
	// In a real scenario, this function would be provided.
	return 0
}
