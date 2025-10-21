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
	root := &TreeNode{5, nil, nil}
	root.Left = &TreeNode{3, nil, nil}
	root.Right = &TreeNode{6, nil, nil}
	root.Left.Left = &TreeNode{2, nil, nil}
	root.Left.Right = &TreeNode{4, nil, nil}
	root.Right.Right = &TreeNode{7, nil, nil}

	key := 3
	newRoot := deleteNode(root, key)
	fmt.Println(newRoot) // Output the new tree structure after deletion
}

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}

	if key < root.Val {
		root.Left = deleteNode(root.Left, key)
	} else if key > root.Val {
		root.Right = deleteNode(root.Right, key)
	} else {
		if root.Left == nil {
			return root.Right
		} else if root.Right == nil {
			return root.Left
		} else {
			minLargerNode := root.Right
			for minLargerNode.Left != nil {
				minLargerNode = minLargerNode.Left
			}
			root.Val = minLargerNode.Val
			root.Right = deleteNode(root.Right, minLargerNode.Val)
		}
	}
	return root
}
