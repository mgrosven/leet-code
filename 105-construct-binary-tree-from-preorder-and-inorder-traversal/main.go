package main

func main() {

	p := []int{}
	i := []int{}

	buildTree(p, i)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	rootVal := preorder[0]
	root := &TreeNode{Val: rootVal}

	rootIndex := 0
	for i, val := range inorder {
		if val == rootVal {
			rootIndex = i
			break
		}
	}

	leftInOrder := inorder[:rootIndex]
	rightInOrder := inorder[rootIndex+1:]

	leftPreOrder := preorder[1 : 1+rootIndex]
	rightPreOrder := preorder[1+rootIndex:]

	root.Left = buildTree(leftPreOrder, leftInOrder)
	root.Right = buildTree(rightPreOrder, rightInOrder)
	return root
}
