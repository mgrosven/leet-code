package main

import (
	"container/heap"
	"fmt"
)

func main() {
	nums := []int{1, 1, 1, 2, 2, 3}
	k := 1
	r := topKFrequent(nums, k)
	fmt.Println(r)
}

func topKFrequent(nums []int, k int) []int {
	counts := make(map[int]int)

	for _, v := range nums {
		counts[v]++
	}

	pq := &PriorityQueue{}
	heap.Init(pq)

	for key, value := range counts {
		heap.Push(pq, &Item{value: key, frequency: value})
		if pq.Len() > k {
			heap.Pop(pq)
		}
	}

	var result []int

	for pq.Len() > 0 {
		item := heap.Pop(pq).(*Item)
		result = append(result, item.value)
	}

	return result
}

type Item struct {
	value     int
	frequency int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].frequency < pq[j].frequency
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(*Item))
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item

}
