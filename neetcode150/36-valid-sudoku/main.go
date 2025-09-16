package main

import "fmt"

func main() {
	matrix := [][]byte{
		{'8', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}
	r := isValidSudoku(matrix)
	if r {
		fmt.Println("Valid")
	} else {
		fmt.Println("Invalid")
	}
}

func isValidSudoku(board [][]byte) bool {
	var rows, columns, boxes [9][9]bool
	for i, r := range board {
		for j, e := range r {
			if e != '.' {
				v := e - '0' - 1
				b := (i/3)*3 + (j / 3)
				if rows[i][v] || columns[j][v] || boxes[b][v] {
					return false
				}
				rows[i][v], columns[j][v], boxes[b][v] = true, true, true
			}
		}
	}
	return true
}
