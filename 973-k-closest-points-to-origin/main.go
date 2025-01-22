package main

import (
	"container/heap"
	"fmt"
	"math"
)

func main() {

	points := [][]int{{3, 3}, {5, -1}, {-2, 4}}
	k := 2
	fmt.Println(kClosest(points, k))

}

func kClosest(points [][]int, k int) [][]int {
	h := &IntHeap{}
	heap.Init(h)

	for _, p := range points {
		d := math.Sqrt(math.Pow(float64(p[0]), 2) + math.Pow(float64(p[1]), 2))

		heap.Push(h, Point{
			Loc:  p,
			Dist: d,
		})
	}

	results := [][]int{}
	for i := 0; i < k; i++ {
		p := heap.Pop(h).(Point)
		results = append(results, p.Loc)
	}
	return results
}

type Point struct {
	Loc  []int
	Dist float64
}

type IntHeap []Point

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i].Dist < h[j].Dist }
func (h IntHeap) Swap(i, j int) {
	if h.Len() > 1 {
		h[i], h[j] = h[j], h[i]
	}
}

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(Point))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
