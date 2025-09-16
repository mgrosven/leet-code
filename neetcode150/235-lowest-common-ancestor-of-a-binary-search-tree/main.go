package main

func main() {

	// l1 := makeList([]int{1, 2, 3, 4})

	lowestCommonAncestor(&TreeNode{}, &TreeNode{})
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root != nil {

		if p.Val < root.Val && q.Val < root.Val {
			return lowestCommonAncestor(root.Left, p, q)
		} else if p.Val > root.Val && q.Val > root.Val {
			return lowestCommonAncestor(root.Right, p, q)
		}
	}
	return root
}
