package main

func main() {

	// l1 := makeList([]int{1, 2, 3, 4})

	maxDepth(&TreeNode{})
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	return checkDepth(root, 1)

}

func checkDepth(root *TreeNode, depth int) int {
	if root != nil {
		return max(checkDepth(root.Left, depth+1), checkDepth(root.Right, depth+1))
	}
	return depth - 1
}
