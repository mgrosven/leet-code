package main

import (
	"fmt"
)

func main() {

	l1 := makeList([]int{1, 2, 3, 4, 5})
	n := 2

	o := removeNthFromEnd(l1, n)
	printList(o)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func makeList(nums []int) *ListNode {
	l := &ListNode{}
	next := l

	for i, n := range nums {
		next.Val = n
		if i != len(nums)-1 {
			next.Next = &ListNode{}
			next = next.Next
		}
	}
	return l
}

func printList(head *ListNode) {
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

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	p1 := head
	for i := 0; i < n; i++ {
		p1 = p1.Next
	}

	var prev *ListNode
	p2 := head
	for p1 != nil {
		p1 = p1.Next
		prev = p2
		p2 = p2.Next
	}

	if prev == nil {
		return head.Next
	}

	prev.Next = p2.Next
	return head
}
