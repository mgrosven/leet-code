package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{1, nil, nil}
	root.Left = &TreeNode{2, nil, nil}
	root.Right = &TreeNode{3, nil, nil}
	root.Left.Left = &TreeNode{4, nil, nil}
	root.Left.Right = &TreeNode{5, nil, nil}
	root.Right.Right = &TreeNode{6, nil, nil}

	fmt.Println(maxLevelSum(root)) // Output: 3
}

func maxLevelSum(root *TreeNode) int {
	if root == nil {
		return 0
	}

	maxSum := root.Val
	maxLevel := 1
	currentLevel := 1

	queue := []*TreeNode{root}

	for len(queue) > 0 {
		levelSize := len(queue)
		currentSum := 0

		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]

			currentSum += node.Val

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		if currentSum > maxSum {
			maxSum = currentSum
			maxLevel = currentLevel
		}

		currentLevel++
	}

	return maxLevel
}
