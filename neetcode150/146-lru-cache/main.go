package main

type Node struct {
	Key, Value int
	Prev, Next *Node
}

type LRUCache struct {
	Capacity int
	Cache    map[int]*Node
	Head     *Node
	Tail     *Node
}

func Constructor(capacity int) LRUCache {
	head := &Node{}
	tail := &Node{}
	head.Next = tail
	tail.Prev = head
	return LRUCache{
		Capacity: capacity,
		Cache:    make(map[int]*Node),
		Head:     head,
		Tail:     tail,
	}
}

func (this *LRUCache) Get(key int) int {
	if node, exists := this.Cache[key]; exists {
		this.moveToFront(node)
		return node.Value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if node, exists := this.Cache[key]; exists {
		node.Value = value
		this.moveToFront(node)
	} else {
		newNode := &Node{Key: key, Value: value}
		this.Cache[key] = newNode
		this.addNode(newNode)

		if len(this.Cache) > this.Capacity {
			tail := this.popTail()
			delete(this.Cache, tail.Key)
		}
	}
}

func (this *LRUCache) addNode(node *Node) {
	node.Prev = this.Head
	node.Next = this.Head.Next
	this.Head.Next.Prev = node
	this.Head.Next = node
}

func (this *LRUCache) removeNode(node *Node) {
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
}

func (this *LRUCache) moveToFront(node *Node) {
	this.removeNode(node)
	this.addNode(node)
}

func (this *LRUCache) popTail() *Node {
	node := this.Tail.Prev
	this.removeNode(node)
	return node
}

func main() {
	lRUCache := Constructor(2)
	lRUCache.Put(1, 1) // cache is {1=1}
	lRUCache.Put(2, 2) // cache is {1=1, 2=2}
	lRUCache.Get(1)    // return 1
	lRUCache.Put(3, 3) // LRU key was 2, evicts key 2, cache is {1=1, 3=3}
	lRUCache.Get(2)    // returns -1 (not found)
	lRUCache.Put(4, 4) // LRU key was 1, evicts key 1, cache is {4=4, 3=3}
	lRUCache.Get(1)    // return -1 (not found)
	lRUCache.Get(3)    // return 3
	lRUCache.Get(4)    // return 4
}
