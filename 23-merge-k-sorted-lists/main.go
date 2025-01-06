package main

import (
	"container/heap"
	"fmt"
)

func main() {

	l1 := []*ListNode{
		makeList([]int{1, 4, 5}),
		makeList([]int{1, 3, 4}),
		makeList([]int{2, 6}),
	}

	l := mergeKLists(l1)
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

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	h := BuildNewMinHeap(lists)

	if h.IsEmpty() {
		return nil
	}

	head, ok := heap.Pop(h).(*ListNode)
	if !ok {
		return nil
	}

	list := head
	for !h.IsEmpty() {
		list.Next, ok = heap.Pop(h).(*ListNode)
		list = list.Next
		if !ok {
			return nil
		}
	}
	return head
}

type MinHeap struct {
	items []*ListNode
}

func NewMinHeap() *MinHeap {
	h := &MinHeap{
		items: make([]*ListNode, 0),
	}
	heap.Init(h)
	return h
}

func (h *MinHeap) Len() int {
	return len(h.items)
}

func (h *MinHeap) Less(i, j int) bool {
	return h.items[i].Val < h.items[j].Val
}

func (h *MinHeap) Swap(i, j int) {
	h.items[i], h.items[j] = h.items[j], h.items[i]
}

func (h *MinHeap) Push(x any) {
	item, ok := x.(*ListNode)
	if !ok {
		return
	}
	h.items = append(h.items, item)
}

func (h *MinHeap) Pop() (top any) {
	size := len(h.items)
	if size == 0 {
		return nil
	}

	top, h.items = h.items[size-1], h.items[:size-1]
	return top
}

func (h *MinHeap) IsEmpty() bool {
	return len(h.items) == 0
}

func BuildNewMinHeap(lists []*ListNode) *MinHeap {
	h := NewMinHeap()
	for _, list := range lists {
		if list == nil {
			continue
		}
		var prev, cur *ListNode = nil, list
		for cur != nil {
			prev, cur = cur, cur.Next
			prev.Next = nil
			heap.Push(h, prev)
		}
	}
	return h
}
