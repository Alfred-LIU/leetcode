package DoubleLinkedList

/*
Design a data structure that follows the constraints of a Least Recently Used (LRU) cache.

Implement the LRUCache class:

LRUCache(int capacity) Initialize the LRU cache with positive size capacity.
int get(int key) Return the value of the key if the key exists, otherwise return -1.
void put(int key, int value) Update the value of the key if the key exists. Otherwise, add the key-value pair to the cache. If the number of keys exceeds the capacity from this operation, evict the least recently used key.
The functions get and put must each run in O(1) average time complexity.
*/
type Node struct {
	key  int
	val  int
	pre  *Node
	next *Node
}

type LRUCache struct {
	s        map[int]*Node
	capacity int
	tail     *Node
	head     *Node
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		s:        make(map[int]*Node),
		capacity: capacity,
	}
}

func (this *LRUCache) Get(key int) int {
	if val, ok := this.s[key]; ok {
		this.remove(val)
		this.add(val)
		return val.val
	} else {
		return -1
	}
}

func (this *LRUCache) Put(key int, value int) {
	if val, ok := this.s[key]; ok {
		this.remove(val)
		val.val = value
		this.add(val)
		return
	} else {
		node := &Node{
			key: key,
			val: value,
		}
		this.s[key] = node
		this.add(node)
		if len(this.s) > this.capacity {
			delete(this.s, this.tail.key)
			this.remove(this.tail)
		}
	}
}

func (this *LRUCache) add(node *Node) {
	node.pre = nil
	node.next = this.head
	if this.head != nil {
		this.head.pre = node
	}
	this.head = node
	if this.tail == nil {
		this.tail = node
	}
}

func (this *LRUCache) remove(node *Node) {
	if node != this.head {
		node.pre.next = node.next
	} else {
		this.head = node.next
	}

	if node != this.tail {
		node.next.pre = node.pre
	} else {
		this.tail = node.pre
	}
}
