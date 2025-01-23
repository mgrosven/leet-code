package main

import (
	"container/heap"
	"fmt"
)

func main() {

	taskList := []byte{'A', 'A', 'A', 'B', 'B', 'B'}
	n := 2
	fmt.Println(leastInterval(taskList, n))
}

func leastInterval(tasks []byte, n int) int {
	if n == 0 {
		return len(tasks)
	}
	m := make(map[byte]int)

	for _, t := range tasks {
		m[t]++
	}

	h := &TaskHeap{}
	for k, v := range m {
		t := Task{
			Name:  k,
			Count: v,
		}
		h.Push(t)
	}
	heap.Init(h)

	time := 0
	for h.Len() > 0 {
		temp := []Task{}
		for i := 0; i < n+1; i++ {
			t := heap.Pop(h).(Task)
			if t.Count > 1 {
				t.Count--
				temp = append(temp, t)
			}
			time++
			if h.Len() == 0 && len(temp) == 0 {
				break
			}
		}

		for _, t := range temp {
			heap.Push(h, t)
		}
	}
	return time
}

type Task struct {
	Name  byte
	Count int
	Cycle int
}

type TaskHeap []Task

func (h TaskHeap) Len() int           { return len(h) }
func (h TaskHeap) Less(i, j int) bool { return h[i].Count > h[j].Count }
func (h TaskHeap) Swap(i, j int) {
	if h.Len() > 1 {
		h[i], h[j] = h[j], h[i]
	}
}

func (h *TaskHeap) Push(x interface{}) {
	*h = append(*h, x.(Task))
}

func (h *TaskHeap) Pop() interface{} {
	if h.Len() > 0 {
		current := *h
		n := len(current)
		x := current[n-1]
		*h = current[:n-1]
		return x
	}
	return Task{}
}
