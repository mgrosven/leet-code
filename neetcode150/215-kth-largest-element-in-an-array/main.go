package main

import (
	"container/heap"
	"fmt"
)

func main() {
	nums := []int{3, 2, 3, 1, 2, 4, 5, 5, 6}
	k := 4
	fmt.Println(findKthLargest(nums, k))
}

func findKthLargest(nums []int, k int) int {
	h := &IntHeap{}
	heap.Init(h)
	for _, v := range nums {
		heap.Push(h, v)
	}

	var result int
	for i := 0; i < k; i++ {
		result = heap.Pop(h).(int)
	}
	return result
}

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntHeap) Swap(i, j int) {
	if h.Len() > 1 {
		h[i], h[j] = h[j], h[i]
	}
}

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
