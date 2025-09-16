package main

import (
	"container/heap"
	"fmt"
)

func main() {
	s := []int{2, 2}

	fmt.Println(lastStoneWeight(s))

}

func lastStoneWeight(stones []int) int {

	h := &IntHeap{}
	heap.Init(h)
	for _, v := range stones {
		heap.Push(h, v)
	}

	fmt.Println(*h)

	for h.Len() > 1 {
		y := heap.Pop(h).(int)
		x := heap.Pop(h).(int)

		if x != y {
			heap.Push(h, y-x)
		}
	}
	if h.Len() > 0 {
		return heap.Pop(h).(int)
	}
	return 0
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
