package custom

import (
	"fmt"
	"testing"
)

func TestMinHeap(t *testing.T) {
	arr := []int{3, 6, 1}
	// heap.Init(arr)
	// fmt.Printf("Min value: %v\n", heap.Pop(arr))
	// heap.Push(arr, 5)
	// for arr.Len() > 0 {
	// 	fmt.Printf("Next min: %v\n", heap.Pop(arr))
	// }
	h := InitMinHeap(arr)
	fmt.Println(h.Peek())
	fmt.Println(h.GetMin())
	fmt.Println(h.GetMin())
}

func TestMaxHeap(t *testing.T) {
	arr := []int{1, 2, 6}
	h := InitMaxHeap(arr)

	fmt.Println(h.Peek())
	fmt.Println(h.GetMax())
	fmt.Println(h.GetMax())
}
