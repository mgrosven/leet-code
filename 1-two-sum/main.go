package main

import "fmt"

func main() {
	a := []int{2, 7, 11, 15}
	t := 26
	r := twoSum(a, t)
	fmt.Println(r)
}

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, n := range nums {
		if i == 0 {
			m[target-n] = i + 1
		} else {
			if m[n] != 0 {
				return []int{m[n] - 1, i}
			} else {

				m[target-n] = i + 1
			}
		}
	}
	return []int{0, 0}
}
