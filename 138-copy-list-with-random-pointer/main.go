package main

import (
	"fmt"
)

func main() {

	// l1 := [][]int{{7, nil}, {}}
	// n := 2

	// o := copyRandomList(l1, n)
	// printList(o)
}

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func makeList(nums []int) *Node {
	l := &Node{}
	next := l

	for i, n := range nums {
		next.Val = n
		if i != len(nums)-1 {
			next.Next = &Node{}
			next = next.Next
		}
	}
	return l
}

func printList(head *Node) {
	curr := head
	fmt.Print("[")
	for curr != nil {
		fmt.Print(curr.Val)
		if curr.Next != nil {
			fmt.Print(" -> ")
		}
		curr = curr.Next
	}
	fmt.Print("]")
	fmt.Println()
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}

	nodeMap := make(map[*Node]*Node)

	cur := head
	for cur != nil {
		nodeMap[cur] = &Node{Val: cur.Val}
		cur = cur.Next
	}

	cur = head
	for cur != nil {
		nodeMap[cur].Next = nodeMap[cur.Next]
		nodeMap[cur].Random = nodeMap[cur.Random]
		cur = cur.Next
	}

	return nodeMap[head]
}
