package main

import (
	"fmt"
)

func main() {
	isConnected := [][]int{
		{1, 1, 0},
		{1, 1, 0},
		{0, 0, 1},
	}
	result := findCircleNum(isConnected)
	fmt.Println(result) // Output: 2

	isConnected2 := [][]int{
		{1, 0, 0},
		{0, 1, 0},
		{0, 0, 1},
	}
	result2 := findCircleNum(isConnected2)
	fmt.Println(result2) // Output: 3
}

func findCircleNum(isConnected [][]int) int {
	visited := make([]bool, len(isConnected))

	var visit func(int)
	visit = func(i int) {
		visited[i] = true
		for j := range isConnected {
			if isConnected[i][j] == 1 && !visited[j] {
				visit(j)
			}
		}
	}

	count := 0
	for i := range isConnected {
		if !visited[i] {
			visit(i)
			count++
		}
	}

	return count
}
