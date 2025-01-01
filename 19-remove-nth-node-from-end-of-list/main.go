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
	reversed := reverseList(head)
	cur := reversed
	var prev *ListNode
	i := 1
	for i < n && cur != nil {
		prev = cur
		cur = cur.Next
		i++
	}

	if prev == nil {
		reversed = cur.Next
	} else {
		prev.Next = cur.Next
	}

	return reverseList(reversed)
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
