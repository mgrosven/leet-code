package main

import (
	"fmt"
)

func main() {

	l1 := makeList([]int{1, 2, 3, 4, 5})
	k := 2

	l := reverseKGroup(l1, k)
	printList(l)
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

func reverseKGroup(head *ListNode, k int) *ListNode {
	start := head
	var result, end *ListNode
	for hasKNodes(start, k) {
		// Head of reversed section, next node/start of next section
		h, n := reverseNodes(start, k)
		if result == nil {
			// Track head of new list
			result = h
		} else {
			// End of old list -> start of new list
			end.Next = h
		}
		// end of new list is old start of new list
		end = start
		// start of next section
		start = n
	}
	end.Next = start
	return result
}

func hasKNodes(head *ListNode, k int) bool {
	list, count := head, 0
	for count < k {
		if list != nil {
			list = list.Next
			count++
		} else {
			return false
		}
	}
	return true
}

func reverseNodes(head *ListNode, k int) (*ListNode, *ListNode) {
	var prev *ListNode
	curr := head

	count := 0
	for count < k {
		nextNode := curr.Next
		curr.Next = prev
		prev = curr
		curr = nextNode
		count++
	}
	return prev, curr
}
