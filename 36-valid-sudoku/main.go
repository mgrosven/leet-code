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
	row := make(map[int][9]bool)
	column := make(map[int][9]bool)
	box := make(map[int][9]bool)

	for i, r := range board {
		for j, e := range r {
			if e != '.' {
				v := e - '0' - 1
				b := (i/3)*3 + (j / 3)
				if row[i][v] {
					return false
				}
				if column[j][v] {
					return false
				}
				if box[b][v] {
					return false
				}
				tempRow := row[i]
				tempRow[v] = true
				row[i] = tempRow
				tempColumn := column[j]
				tempColumn[v] = true
				column[j] = tempColumn
				tempBox := box[b]
				tempBox[v] = true
				box[b] = tempBox
			}
		}
	}
	return true
}
