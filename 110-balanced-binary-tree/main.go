package main

import "math"

func main() {

	// l1 := makeList([]int{1, 2, 3, 4})

	isBalanced(&TreeNode{})
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	balanced := true
	var traverse func(*TreeNode) (float64, bool)
	traverse = func(node *TreeNode) (float64, bool) {
		if node == nil {
			return 0, true
		}
		left, lb := traverse(node.Left)
		right, rb := traverse(node.Right)
		balanced = lb && rb && math.Abs(left-right) <= 1
		return max(left, right) + 1, balanced
	}
	traverse(root)
	return balanced
}
