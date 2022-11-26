package PriorityQueue

import (
	"container/heap"
)

/*
Design a max stack data structure that supports the stack operations and supports finding the stack's maximum element.

Implement the MaxStack class:

MaxStack() Initializes the stack object.
void push(int x) Pushes element x onto the stack.
int pop() Removes the element on top of the stack and returns it.
int top() Gets the element on the top of the stack without removing it.
int peekMax() Retrieves the maximum element in the stack without removing it.
int popMax() Retrieves the maximum element in the stack and removes it. If there is more than one maximum element, only remove the top-most one.
You must come up with a solution that supports O(1) for each top call and O(logn) for each other call.

https://leetcode.com/problems/max-stack/
https://pkg.go.dev/container/heap
*/

type element struct {
	stackIndex int
	heapIndex  int
	isDie      bool
	value      int
}

//heap

type MaxHeap []*element

// Len -
func (h MaxHeap) Len() int { return len(h) }

// Less -
func (h MaxHeap) Less(i, j int) bool {
	return h[i].value > h[j].value || h[i].value == h[j].value && h[i].stackIndex > h[j].stackIndex
}

// Swap -
func (h MaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
	h[i].heapIndex = i
	h[j].heapIndex = j
}

// Push -
func (h *MaxHeap) Push(x any) {
	l := len(*h)
	item := x.(*element)
	item.heapIndex = l
	*h = append(*h, item)
}

// Pop -
func (h *MaxHeap) Pop() any {
	old := *h
	l := len(old)
	popItem := old[l-1]
	old[l-1] = nil
	popItem.heapIndex = -1
	*h = old[0 : l-1]
	return popItem
}

// Stack
type stack struct {
	data []*element
	idx  int
	cap  int
}

// newStack -
func newStack(cap int) *stack {
	return &stack{
		data: make([]*element, cap),
		idx:  0,
		cap:  cap,
	}
}

// push -
func (s *stack) push(e *element) {
	if s.idx == s.cap {
		s.grow()
	}

	e.stackIndex = s.idx
	s.data[s.idx] = e
	s.idx++
}

// pop -
func (s *stack) pop() *element {
	s.idx--
	v := s.data[s.idx]
	s.data[s.idx] = nil
	return v
}

// peek
func (s *stack) peek() *element {
	return s.data[s.idx-1]
}

// markDie
func (s *stack) markDie(idx int) {
	if idx >= s.cap {
		return
	}
	s.data[idx].isDie = true
}

// grow -
func (s *stack) grow() {
	newData := make([]*element, s.cap*2)
	copy(newData, s.data)
	s.data = newData
	s.cap = s.cap * 2
}

// start
type MaxStack struct {
	stack *stack
	max   MaxHeap
}

func Constructor() MaxStack {
	return MaxStack{
		stack: newStack(1),
		max:   MaxHeap{},
	}
}

func (ms *MaxStack) Push(x int) {
	el := &element{value: x}
	ms.stack.push(el)
	heap.Push(&ms.max, el)
}

func (ms *MaxStack) Pop() int {
	r := 0
	for {
		el := ms.stack.pop()
		if !el.isDie {
			heap.Remove(&ms.max, el.heapIndex)
			r = el.value
			break
		}
	}

	return r
}

func (ms *MaxStack) Top() int {
	r := 0
	for {
		el := ms.stack.peek()
		if !el.isDie {
			r = el.value
			break
		}
		ms.stack.pop()
	}

	return r
}

func (ms *MaxStack) PeekMax() int {
	return ms.max[0].value
}

func (ms *MaxStack) PopMax() int {
	el := heap.Pop(&ms.max).(*element)
	ms.stack.markDie(el.stackIndex)

	return el.value
}
