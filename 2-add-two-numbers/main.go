package main

import (
	"fmt"
)

func main() {

	l1 := makeList([]int{2, 4, 3})
	l2 := makeList([]int{5, 6, 4})

	o := addTwoNumbers(l1, l2)
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

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	list := head
	carry := 0
	for l1 != nil && l2 != nil {
		sum := l1.Val + l2.Val + carry
		carry = sum / 10
		list.Next = &ListNode{
			Val: sum % 10,
		}
		l1, l2, list = l1.Next, l2.Next, list.Next
	}
	for l1 != nil {
		sum := l1.Val + carry
		carry = sum / 10
		list.Next = &ListNode{
			Val: sum % 10,
		}
		l1, list = l1.Next, list.Next
	}
	for l2 != nil {
		sum := l2.Val + carry
		carry = sum / 10
		list.Next = &ListNode{
			Val: sum % 10,
		}
		l2, list = l2.Next, list.Next
	}
	if carry > 0 {
		list.Next = &ListNode{
			Val: carry,
		}
	}

	return head.Next
}

func getDigitsReversed(n int) []int {
	digits := []int{}
	for n > 0 {
		digits = append(digits, n%10)
		n /= 10
	}
	return digits
}
