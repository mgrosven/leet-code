package main

import (
	"fmt"
)

func main() {

	l := &ListNode{}
	next := l

	nums := []int{1, 2, 3, 4, 5}

	for _, n := range nums {
		next.Val = n
		next.Next = &ListNode{}
		next = next.Next
	}

	next.Next = nil

	o := reverseList(l)
	printList(o)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func printList(head *ListNode) {
	curr := head
	for curr != nil {
		fmt.Print(curr.Val)
		if curr.Next != nil {
			fmt.Print(" -> ")
		}
		curr = curr.Next
	}
	fmt.Println()
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head

	for curr != nil {
		nextNode := curr.Next
		curr.Next = prev
		prev = curr
		curr = nextNode
	}
	return prev
}
