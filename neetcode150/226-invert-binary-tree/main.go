package main

func main() {

	// l1 := makeList([]int{1, 2, 3, 4})

	invertTree(&TreeNode{})
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	root.Left, root.Right = root.Right, root.Left
	if root.Left != nil {
		invertTree(root.Left)
	}
	if root.Right != nil {
		invertTree(root.Right)
	}
	return root
}
