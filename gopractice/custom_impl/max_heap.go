package custom

import "container/heap"

// MaxHeap is heap impl with ints array
type MaxHeap []int

// Functions for heap interface/property

// Len of heap array
func (h MaxHeap) Len() int {
	return len(h)
}

// Less property decides min(<) / max(>) heap
func (h MaxHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

// Swap to swap
func (h MaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

// Push new int
func (h *MaxHeap) Push(val interface{}) {
	*h = append(*h, val.(int))
}

// Pop the min/max int
func (h *MaxHeap) Pop() interface{} {
	n := len(*h)
	res := (*h)[n-1]
	*h = (*h)[:n-1]
	return res
}

// Functions for MaxHeap wrapper

// InitMaxHeap inits a new min heap from arr
func InitMaxHeap(arr []int) *MaxHeap {
	h := MaxHeap(arr)
	heap.Init(&h)
	return &h
}

// GetMax returns max element
func (h *MaxHeap) GetMax() int {
	return heap.Pop(h).(int)
}

// Add an element
func (h *MaxHeap) Add(v int) {
	heap.Push(h, v)
}

// Peek max element
func (h *MaxHeap) Peek() int {
	res := heap.Pop(h).(int)
	heap.Push(h, res)
	return res
}
