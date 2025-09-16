package main

import (
	"fmt"
)

func main() {

	pos := -1
	l1 := makeList([]int{1}, pos)

	o := hasCycle(l1)
	if o {
		fmt.Println("Has Cycle")
	} else {
		fmt.Println("No Cycle")
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func makeList(nums []int, pos int) *ListNode {
	l := &ListNode{}
	next := l

	var tail *ListNode

	for i, n := range nums {
		next.Val = n
		if i == pos {
			tail = next
		}
		if i != len(nums)-1 {
			next.Next = &ListNode{}
			next = next.Next
		}
	}
	next.Next = tail
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

func hasCycle(head *ListNode) bool {
	hashTable := make(map[*ListNode]int)

	for head != nil {
		if _, ok := hashTable[head]; ok {
			return true
		}
		hashTable[head] = 1
		head = head.Next
	}
	return false
}
