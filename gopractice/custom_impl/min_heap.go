package custom

import "container/heap"

// MinHeap is a min heap impl with ints
type MinHeap []int

// Functions for heap interface/property

// Len of heap array
func (h MinHeap) Len() int {
	return len(h)
}

// Less property decides min/max heap
func (h MinHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

// Swap to swap
func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

// Push nw int
func (h *MinHeap) Push(val interface{}) {
	*h = append(*h, val.(int))
}

// Pop the min int
func (h *MinHeap) Pop() interface{} {
	n := len(*h)
	res := (*h)[n-1]
	*h = (*h)[:n-1]
	return res
}

// Functions for MinHeap wrapper

// InitMinHeap inits a new min heap from arr
func InitMinHeap(arr []int) *MinHeap {
	h := MinHeap(arr)
	heap.Init(&h)
	return &h
}

// GetMin returns min element
func (h *MinHeap) GetMin() int {
	return heap.Pop(h).(int)
}

// Add an elemnt
func (h *MinHeap) Add(v int) {
	heap.Push(h, v)
}

// Peek min lemen
func (h *MinHeap) Peek() int{
	res := heap.Pop(h).(int)
	heap.Push(h, res)
	return res
}