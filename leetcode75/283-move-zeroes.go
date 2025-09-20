package main

import "fmt"

func main() {
	input := []int{0, 1, 0, 3, 12}
	moveZeroes(input)
	fmt.Println(input)
}

func moveZeroes(nums []int) {
	insertPos := 0
	for _, num := range nums {
		if num != 0 {
			nums[insertPos] = num
			insertPos++
		}
	}
	for i := insertPos; i < len(nums); i++ {
		nums[i] = 0
	}
}
