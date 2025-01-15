package main

import (
	"fmt"
)

func main() {

	bst := createBST([]interface{}{5, 3, 6, 2, 4, nil, nil, 1})
	k := 3

	r := kthSmallest(bst, k)

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

func kthSmallest(root *TreeNode, k int) int {
	var result int
	var count int

	var inOrder func(node *TreeNode)
	inOrder = func(node *TreeNode) {
		if node == nil || count >= k {
			return
		}

		inOrder(node.Left)

		count++
		if count == k {
			result = node.Val
			return
		}

		inOrder(node.Right)
	}
	inOrder(root)
	return result
}
