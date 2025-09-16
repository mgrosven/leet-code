package main

import "fmt"

func main() {

	// l1 := makeTree([]int{1, 2, 3, 4})

	levelOrder(&TreeNode{})
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// func makeTree(list []int) *TreeNode {
// 	root := &TreeNode{}

// 	for _, v := range list {

// 	}
// }

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	q := []*TreeNode{root}
	result := [][]int{}

	for len(q) > 0 {
		levelSize := len(q)
		level := []int{}

		for i := 0; i < levelSize; i++ {
			e := q[0]
			level = append(level, e.Val)
			if e.Left != nil {
				q = append(q, e.Left)
			}
			if e.Right != nil {
				q = append(q, e.Right)
			}
			q = q[1:]
		}
		result = append(result, level)
	}
	return result
}
