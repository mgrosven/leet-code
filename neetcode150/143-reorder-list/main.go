package main

import (
	"fmt"
)

func main() {

	l1 := makeList([]int{1, 2, 3, 4})

	reorderList(l1)
	printList(l1)
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
	for curr != nil {
		fmt.Print(curr.Val)
		if curr.Next != nil {
			fmt.Print(" -> ")
		}
		curr = curr.Next
	}
	fmt.Println()
}

func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}

	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	second := reverseList(slow.Next)
	slow.Next = nil
	first := head
	for second != nil {
		t1, t2 := first.Next, second.Next
		first.Next = second
		second.Next = t1
		first = t1
		second = t2
	}
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
