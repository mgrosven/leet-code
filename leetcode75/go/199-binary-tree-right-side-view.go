package main

import (
	"fmt"
)

func main() {
	root := &TreeNode{1, nil, nil}
	root.Right = &TreeNode{3, nil, nil}
	root.Left = &TreeNode{2, nil, nil}
	root.Left.Right = &TreeNode{5, nil, nil}
	root.Right.Right = &TreeNode{4, nil, nil}

	result := rightSideView(root)
	fmt.Println(result) // Output: [1, 3, 4]
}

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	result := []int{}
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		queueLen := len(queue)
		for i := 0; i < queueLen; i++ {
			node := queue[0]
			queue = queue[1:]

			if i == queueLen-1 {
				result = append(result, node.Val)
			}

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}
	return result
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
