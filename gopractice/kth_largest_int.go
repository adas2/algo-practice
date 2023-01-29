// Find kth largest element in an int arr
package practice

import (
	custom "adas2.io/practice/custom_impl"
)

// solution using heap

func findKthLargest(nums []int, k int) int {

	// create empty min-heap
	h := custom.InitMinHeap([]int{})

	for _, elem := range nums {
		// add element to heap
		h.Add(elem)
		// if size > k, pop smallest element
		if h.Len() > k {
			h.GetMin()
		}
	}

	// return top of heap
	return h.GetMin()
}

// Complexity, since heap size never exceeds k, each operation in O(logk)
// Total time =  N x O(logk)
