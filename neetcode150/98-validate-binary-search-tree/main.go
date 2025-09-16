package main

import "math"

func main() {

	// l1 := makeTree([]int{1, 2, 3, 4})

	isValidBST(&TreeNode{})
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

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	isValid := true

	var checkNode func(node *TreeNode, min, max int)
	checkNode = func(node *TreeNode, min, max int) {
		if node == nil || !isValid {
			return
		}
		if node.Val <= min || node.Val >= max {
			isValid = false
			return
		}
		checkNode(node.Left, min, node.Val)
		checkNode(node.Right, node.Val, max)
	}
	checkNode(root, math.MinInt, math.MaxInt)
	return isValid
}
