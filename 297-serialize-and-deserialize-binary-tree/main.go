package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	bst := createBST([]interface{}{1, 2, 3, nil, nil, 4, 5})

	// fmt.Println(r)

	ser := Constructor()
	deser := Constructor()
	data := ser.serialize(bst)
	ans := deser.deserialize(data)

	fmt.Println(ans)
}

// createBST builds a BST from an array of values
func createBST(values []interface{}) *TreeNode {
	if len(values) == 0 || values[0] == nil {
		return nil
	}
	root := &TreeNode{Val: values[0].(int)}
	queue := []*TreeNode{root} // Queue for level-order traversal
	idx := 1                   // Pointer to the next value in the input

	for len(queue) > 0 && idx < len(values) {
		// Get the current node from the queue
		current := queue[0]
		queue = queue[1:]

		// Process the left child
		if idx < len(values) && values[idx] != nil {
			current.Left = &TreeNode{Val: values[idx].(int)}
			queue = append(queue, current.Left)
		}
		idx++

		// Process the right child
		if idx < len(values) && values[idx] != nil {
			current.Right = &TreeNode{Val: values[idx].(int)}
			queue = append(queue, current.Right)
		}
		idx++
	}

	return root
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}

	var result []string
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if node == nil {
			result = append(result, "nil")
		} else {
			result = append(result, strconv.Itoa(node.Val))
			queue = append(queue, node.Left, node.Right)
		}
	}

	return strings.Join(result, ",")
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	if data == "" {
		return nil
	}
	values := strings.Split(data, ",")

	rootVal, _ := strconv.Atoi(values[0])
	root := &TreeNode{Val: rootVal}
	queue := []*TreeNode{root}

	idx := 1
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if values[idx] != "nil" {
			leftVal, _ := strconv.Atoi(values[idx])
			current.Left = &TreeNode{Val: leftVal}
			queue = append(queue, current.Left)
		}
		idx++

		if values[idx] != "nil" {
			rightVal, _ := strconv.Atoi(values[idx])
			current.Right = &TreeNode{Val: rightVal}
			queue = append(queue, current.Right)
		}
		idx++
	}

	return root
}
