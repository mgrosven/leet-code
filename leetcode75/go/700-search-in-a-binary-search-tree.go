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
	root := &TreeNode{4, nil, nil}
	root.Left = &TreeNode{2, nil, nil}
	root.Right = &TreeNode{7, nil, nil}
	root.Left.Left = &TreeNode{1, nil, nil}
	root.Left.Right = &TreeNode{3, nil, nil}

	val := 2
	result := searchBST(root, val)
	if result != nil {
		fmt.Println(result.Val) // Output: 2
	} else {
		fmt.Println("Value not found")
	}
}

func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == val {
		return root
	} else if val < root.Val {
		return searchBST(root.Left, val)
	} else {
		return searchBST(root.Right, val)
	}
}
