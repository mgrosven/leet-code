package main

import (
	"fmt"
)

func main() {
	node := &ListNode{1, nil}
	node.Next = &ListNode{3, nil}
	node.Next.Next = &ListNode{4, nil}
	node.Next.Next.Next = &ListNode{7, nil}
	node.Next.Next.Next.Next = &ListNode{1, nil}
	node.Next.Next.Next.Next.Next = &ListNode{2, nil}
	node.Next.Next.Next.Next.Next.Next = &ListNode{6, nil}

	printList(node)
	result := deleteMiddle(node)
	printList(result)
}

func deleteMiddle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}

	slow, fast := head, head
	var prev *ListNode

	for fast != nil && fast.Next != nil {
		prev = slow
		slow = slow.Next
		fast = fast.Next.Next
	}

	prev.Next = slow.Next
	return head
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func printList(head *ListNode) {
	for head != nil {
		fmt.Printf("%d -> ", head.Val)
		head = head.Next
	}
	fmt.Println("nil")
}
