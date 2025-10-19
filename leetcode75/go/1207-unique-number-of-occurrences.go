package main

import (
	"fmt"
)

func main() {
	arr := []int{1, 2, 2, 1, 1, 3}
	result := uniqueOccurrences(arr)
	fmt.Println(result) // Output: true

	arr = []int{1, 2}
	result = uniqueOccurrences(arr)
	fmt.Println(result) // Output: false

	arr = []int{-3, 0, 1, -3, 1, 1, 1, -3, 10, 0}
	result = uniqueOccurrences(arr)
	fmt.Println(result) // Output: true
}

func uniqueOccurrences(arr []int) bool {
	frequencyMap := make(map[int]int)
	for _, num := range arr {
		frequencyMap[num]++
	}

	occurrenceSet := make(map[int]struct{})
	for _, freq := range frequencyMap {
		if _, exists := occurrenceSet[freq]; exists {
			return false
		}
		occurrenceSet[freq] = struct{}{}
	}

	return true
}
