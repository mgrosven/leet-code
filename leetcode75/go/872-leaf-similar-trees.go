package main

import (
	"fmt"
)

func main() {
	root1 := &TreeNode{1, nil, nil}
	root1.Left = &TreeNode{2, nil, nil}
	root1.Right = &TreeNode{3, nil, nil}
	root1.Left.Left = &TreeNode{4, nil, nil}
	root1.Left.Right = &TreeNode{5, nil, nil}
	root1.Right.Right = &TreeNode{6, nil, nil}

	root2 := &TreeNode{1, nil, nil}
	root2.Left = &TreeNode{2, nil, nil}
	root2.Right = &TreeNode{3, nil, nil}
	root2.Left.Left = &TreeNode{4, nil, nil}
	root2.Left.Right = &TreeNode{5, nil, nil}
	root2.Right.Right = &TreeNode{6, nil, nil}

	fmt.Println("Leaf Similar:", leafSimilar(root1, root2))
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	leaves1 := []int{}
	leaves2 := []int{}

	var dfs func(node *TreeNode, leaves *[]int)
	dfs = func(node *TreeNode, leaves *[]int) {
		if node == nil {
			return
		}
		if node.Left == nil && node.Right == nil {
			*leaves = append(*leaves, node.Val)
			return
		}
		dfs(node.Left, leaves)
		dfs(node.Right, leaves)
	}

	dfs(root1, &leaves1)
	dfs(root2, &leaves2)

	if len(leaves1) != len(leaves2) {
		return false
	}

	for i := range leaves1 {
		if leaves1[i] != leaves2[i] {
			return false
		}
	}

	return true
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
