package main

import (
	"fmt"
)

func main() {
	m := [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}
	t := 3
	o := searchMatrix(m, t)
	fmt.Println(o)
}

func searchMatrix(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	left, right := 0, m*n-1
	for left <= right {
		mid := (left + right) / 2
		row, column := mid/n, mid%n
		v := matrix[row][column]
		if v == target {
			return true
		} else if v < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return false
}
