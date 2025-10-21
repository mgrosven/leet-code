package main

import (
	"fmt"
)

func main() {
	head := &ListNode{1, nil}
	head.Next = &ListNode{2, nil}
	head.Next.Next = &ListNode{3, nil}
	head.Next.Next.Next = &ListNode{4, nil}
	head.Next.Next.Next.Next = &ListNode{5, nil}

	reversedHead := reverseList(head)

	for node := reversedHead; node != nil; node = node.Next {
		fmt.Println(node.Val)
	}
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	current := head

	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}

	return prev
}

type ListNode struct {
	Val  int
	Next *ListNode
}
