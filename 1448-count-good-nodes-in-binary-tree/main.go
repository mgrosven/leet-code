package main

func main() {

	// l1 := makeTree([]int{1, 2, 3, 4})

	goodNodes(&TreeNode{})
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// func makeTree(list []int) *TreeNode {
// 	root := &TreeNode{}

// 	for _, v := range list {

// 	}
// }

func goodNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}

	total := 1
	var checkNode func(node *TreeNode, checkValue int)
	checkNode = func(node *TreeNode, checkValue int) {
		if node != nil {
			if node.Val >= checkValue {
				checkValue = node.Val
				total++
			}
			checkNode(node.Left, checkValue)
			checkNode(node.Right, checkValue)
		}
	}

	checkNode(root.Left, root.Val)
	checkNode(root.Right, root.Val)

	return total
}
