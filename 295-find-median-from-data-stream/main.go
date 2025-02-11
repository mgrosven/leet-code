package main

import (
    "container/heap"
    "fmt"
)

func main() {
    obj := Constructor()
    obj.AddNum(1)
    obj.AddNum(2)
   fmt.Printf("%.1f\n", obj.FindMedian())
   obj.AddNum(3)
   fmt.Printf("%.1f\n", obj.FindMedian())
}

type MedianFinder struct {
    minHeap *MinHeap // Bigger Half
    maxHeap *MaxHeap // Smaller Half
}


func Constructor() MedianFinder {
    return MedianFinder{
        minHeap: &MinHeap{}, 
        maxHeap: &MaxHeap{},
    }
}


func (this *MedianFinder) AddNum(num int)  {
    heap.Push(this.maxHeap, num)
    heap.Push(this.minHeap, heap.Pop(this.maxHeap))

    if this.maxHeap.Len() < this.minHeap.Len() {
        heap.Push(this.maxHeap, heap.Pop(this.minHeap))
    }    
}

func (this *MedianFinder) FindMedian() float64 {
   if this.maxHeap.Len() > this.minHeap.Len() {
       return float64((*this.maxHeap)[0])
   }
   return float64(((*this.maxHeap)[0] + (*this.minHeap)[0])) / 2.0
}

type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
    *h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
    last := (*h)[len(*h)-1]
    *h = (*h)[:len(*h)-1]
    return last
}

type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x interface{}) {
    *h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() interface{} {
    last := (*h)[len(*h)-1]
    *h = (*h)[:len(*h)-1]
    return last
}
