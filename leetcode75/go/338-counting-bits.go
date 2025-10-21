package main

import (
	"fmt"
)

func main() {
	n := 5
	result := countBits(n)
	fmt.Println(result) // Output: [0 1 1 2 1 2]

	n = 2
	result = countBits(n)
	fmt.Println(result) // Output: [0 1 1]

	n = 0
	result = countBits(n)
	fmt.Println(result) // Output: [0]

	n = 10
	result = countBits(n)
	fmt.Println(result) // Output: [0 1 1 2 1 2 2 3 1 2 2]
}

func countBits(n int) []int {
	result := make([]int, n+1)
	for i := 1; i <= n; i++ {
		result[i] = result[i>>1] + (i & 1)
	}
	return result
}
