package practice

import (
	"fmt"
	"math"
)

// Implement a stack with find_max() method that returns the largest (int) value in the stack
// Expected time O(1) for find_max()
// How about a 0(1) time and O(1) space solution? Check: Problems.md

// MaxStack is a simple array based stack
type MaxStack struct {
	arr    []int // array of values
	ptr    int   // stack top index
	limit  int   // max capacity of the stack
	max    []int // fields for calculating max_value
	maxPtr int   // stack ptr for max stack
}

// GetStackHandle initializes return handle of MaxStack
func GetStackHandle(limit int) *MaxStack {
	st := &MaxStack{
		limit: limit,
		// declare empty array with a max capacity
		arr:    make([]int, 0, limit),
		ptr:    -1,
		max:    make([]int, 0, limit),
		maxPtr: -1,
	}

	return st
}

// Push handles
func (st *MaxStack) Push(val int) error {
	if len(st.arr) >= st.limit {
		return fmt.Errorf("StackFull")
	}
	// add the new value
	st.arr = append(st.arr, val)
	st.ptr++

	// for max array, insert element if largest value of stack empty
	if st.maxPtr == -1 || (st.max[st.maxPtr] < val) {
		st.max = append(st.max, val)
		st.maxPtr++
	}

	return nil

}

// Pop handles
// Note that instead of memory deletion the stPtr is changed
func (st *MaxStack) Pop() error {
	if st.ptr == -1 {
		return fmt.Errorf("StackEmpty")
	}
	// delete element if present in max array
	// note push/pop order is preserved between the stacks
	// hence just peek in the max array
	if st.arr[st.ptr] == st.max[st.maxPtr] {
		st.maxPtr--
	}

	// delete the value from main array by changing stPtr
	st.ptr--

	return nil
}

// Peek handles
func (st *MaxStack) Peek() int {
	if st.ptr == -1 {
		return math.MinInt32
	}
	return st.arr[st.ptr]
}

// IsEmpty handles
func (st *MaxStack) IsEmpty() bool {
	var ret bool = false
	if st.ptr == -1 {
		ret = true
	}

	return ret
}

// FindMax return largest value existent in the stack
func (st *MaxStack) FindMax() int {
	// this check ensures max array is non empty
	if st.ptr == -1 {
		return math.MinInt32
	}
	// return peek of max array
	return st.max[st.maxPtr]
}
