package main

import "fmt"

func main() {

	kthLargest := Constructor(3, []int{5, -1})

	fmt.Println(kthLargest.Add(2))
	fmt.Println(kthLargest.Add(1))
	fmt.Println(kthLargest.Add(-1))
	fmt.Println(kthLargest.Add(3))
	fmt.Println(kthLargest.Add(4))

}

type KthLargest struct {
	K    int
	Nums []int
}

func Constructor(k int, nums []int) KthLargest {
	h := KthLargest{
		K:    k,
		Nums: []int{},
	}

	for _, v := range nums {
		h.insert(v)
	}

	return h
}

func (this *KthLargest) Add(val int) int {
	this.insert(val)

	return this.peak()
}

func (this *KthLargest) heapifyUp(index int) {
	parent := (index - 1) / 2

	for index > 0 && this.Nums[index] < this.Nums[parent] {
		this.Nums[index], this.Nums[parent] = this.Nums[parent], this.Nums[index]
		index = parent
		parent = (index - 1) / 2
	}
}

func (this *KthLargest) heapifyDown(index int) {
	lastIndex := len(this.Nums) - 1
	leftChild := 2*index + 1
	rightChild := 2*index + 2
	smallest := index

	if leftChild <= lastIndex && this.Nums[leftChild] < this.Nums[smallest] {
		smallest = leftChild
	}

	if rightChild <= lastIndex && this.Nums[rightChild] < this.Nums[smallest] {
		smallest = rightChild
	}

	if smallest != index {
		this.Nums[index], this.Nums[smallest] = this.Nums[smallest], this.Nums[index]
		this.heapifyDown(smallest)
	}
}

func (this *KthLargest) insert(val int) {
	if len(this.Nums) < this.K {
		this.Nums = append(this.Nums, val)
		this.heapifyUp(len(this.Nums) - 1)
	} else {
		if val > this.peak() {
			this.extract()
			this.Nums = append(this.Nums, val)
			this.heapifyUp(len(this.Nums) - 1)
		}
	}
}

func (this *KthLargest) extract() int {
	if len(this.Nums) == 0 {
		panic("Heap is empty")
	}

	root := this.Nums[0]
	this.Nums[0] = this.Nums[len(this.Nums)-1]
	this.Nums = this.Nums[:len(this.Nums)-1]
	this.heapifyDown(0)
	return root
}

func (this *KthLargest) peak() int {
	if len(this.Nums) == 0 {
		panic("Heap is empty")
	}
	return this.Nums[0]
}
