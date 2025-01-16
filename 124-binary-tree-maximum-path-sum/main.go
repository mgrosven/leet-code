package main

import (
	"fmt"
	"math"
)

func main() {

	bst := createBST([]interface{}{-3})
	r := maxPathSum(bst)

	fmt.Println(r)
}

// createBST builds a BST from an array of values
func createBST(values []interface{}) *TreeNode {
	if len(values) == 0 || values[0] == nil {
		return nil
	}
	root := &TreeNode{Val: values[0].(int)}
	queue := []*TreeNode{root} // Queue for level-order traversal
	idx := 1                   // Pointer to the next value in the input

	for len(queue) > 0 && idx < len(values) {
		// Get the current node from the queue
		current := queue[0]
		queue = queue[1:]

		// Process the left child
		if idx < len(values) && values[idx] != nil {
			current.Left = &TreeNode{Val: values[idx].(int)}
			queue = append(queue, current.Left)
		}
		idx++

		// Process the right child
		if idx < len(values) && values[idx] != nil {
			current.Right = &TreeNode{Val: values[idx].(int)}
			queue = append(queue, current.Right)
		}
		idx++
	}

	return root
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	maxSum := math.MinInt

	var maxGain func(node *TreeNode) int
	maxGain = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		leftGain := max(maxGain(node.Left), 0)
		rightGain := max(maxGain(node.Right), 0)

		currentSum := node.Val + leftGain + rightGain

		maxSum = max(maxSum, currentSum)
		return node.Val + max(leftGain, rightGain)
	}
	maxGain(root)
	return maxSum
}
