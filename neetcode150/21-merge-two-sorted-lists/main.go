package main

import (
	"fmt"
)

func main() {

	l1 := makeList([]int{1, 2, 4})
	l2 := makeList([]int{1, 3, 4})

	o := mergeTwoLists(l1, l2)
	printList(o)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func makeList(nums []int) *ListNode {
	l := &ListNode{}
	next := l

	for _, n := range nums {
		next.Val = n
		next.Next = &ListNode{}
		next = next.Next
	}
	next.Next = nil
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

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	head := &ListNode{}
	cur := head

	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			cur.Next = list1
			list1 = list1.Next
		} else {
			cur.Next = list2
			list2 = list2.Next
		}
		cur = cur.Next
	}

	if list1 != nil {
		cur.Next = list1
	}
	if list2 != nil {
		cur.Next = list2
	}

	return head.Next
}
