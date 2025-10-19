package main

import "fmt"

func main() {
	candies := []int{2, 3, 5, 1, 3}
	extraCandies := 3
	result := kidsWithCandies(candies, extraCandies)
	fmt.Println(result)
}

func kidsWithCandies(candies []int, extraCandies int) []bool {
	largest := 0
	for _, v := range candies {
		if v > largest {
			largest = v
		}
	}

	result := make([]bool, len(candies))
	for i, v := range candies {
		if v+extraCandies >= largest {
			result[i] = true
		} else {
			result[i] = false
		}
	}
	return result
}
