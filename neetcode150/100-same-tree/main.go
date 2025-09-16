package main

func main() {

	// l1 := makeList([]int{1, 2, 3, 4})

	isSameTree(&TreeNode{}, &TreeNode{})
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p != nil && q != nil {
		return p.Val == q.Val && isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
	}
	return p == nil && q == nil
}
