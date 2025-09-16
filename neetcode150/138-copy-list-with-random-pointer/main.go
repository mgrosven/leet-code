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

	cur := head
	for cur != nil {
		copy := &Node{Val: cur.Val, Next: cur.Next}
		cur.Next = copy
		cur = copy.Next
	}

	cur = head
	for cur != nil {
		if cur.Random != nil {
			cur.Next.Random = cur.Random.Next
		}
		cur = cur.Next.Next
	}

	cur = head
	newHead := head.Next
	for cur != nil {
		copy := cur.Next
		cur.Next = copy.Next
		if copy.Next != nil {
			copy.Next = copy.Next.Next
		}
		cur = cur.Next
	}

	return newHead
}
