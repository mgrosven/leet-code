package main

import "fmt"

func main() {
	s := []int{1, 2, 3}

	fmt.Println(subsets(s))
}

func subsets(nums []int) [][]int {
	res := [][]int{{}}
	for _, n := range nums {
		for _, r := range res {
			res = append(res, append([]int{n}, r...))
		}
	}
	return res
}
