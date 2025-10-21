package main

import (
	"fmt"
)

func main() {
	root := &TreeNode{10, nil, nil}
	root.Left = &TreeNode{5, nil, nil}
	root.Right = &TreeNode{-3, nil, nil}
	root.Left.Left = &TreeNode{3, nil, nil}
	root.Left.Right = &TreeNode{2, nil, nil}
	root.Right.Right = &TreeNode{11, nil, nil}
	root.Left.Left.Left = &TreeNode{3, nil, nil}
	root.Left.Left.Right = &TreeNode{-2, nil, nil}
	root.Left.Right.Right = &TreeNode{1, nil, nil}

	targetSum := 8
	result := pathSum(root, targetSum)
	fmt.Println(result) // Output: 3
}

func pathSum(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}
	var checkPath func(node *TreeNode, currentSum int) int
	checkPath = func(node *TreeNode, currentSum int) int {
		if node == nil {
			return 0
		}
		currentSum -= node.Val
		count := 0
		if currentSum == 0 {
			count++
		}
		count += checkPath(node.Left, currentSum)
		count += checkPath(node.Right, currentSum)
		return count
	}

	return checkPath(root, targetSum) + pathSum(root.Left, targetSum) + pathSum(root.Right, targetSum)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
