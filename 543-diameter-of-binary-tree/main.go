package main

func main() {

	// l1 := makeList([]int{1, 2, 3, 4})

	diameterOfBinaryTree(&TreeNode{})
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	diam := 0
	var traverse func(*TreeNode) int
	traverse = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left, right := traverse(node.Left), traverse(node.Right)
		diam = max(diam, left+right)
		return max(left, right) + 1
	}
	traverse(root)
	return diam
}
