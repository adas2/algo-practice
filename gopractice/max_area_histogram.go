package practice

import custom "adas2.io/practice/custom_impl"

func largestRectangleArea(heights []int) int {
	maxArea := 0
	// use a stack
	st := custom.Sstack{}
	// push a -1 as index to mark the beginning of the stack
	st.Push(-1)
	// traverse the histogram
	index := 0
	for index < len(heights) {
		// ht of top of stack index
		tpIdx := st.Peek()
		// push idx is heights are larger than top of stack
		if tpIdx == -1 || heights[index] >= heights[tpIdx] {
			st.Push(index)
			index++
		} else { // smaller than top index
			// pop top
			tpIdx, _ = st.Pop()
			currHt := heights[tpIdx]
			// calc area with this as min height
			area := currHt * (index - st.Peek() - 1)
			// check max
			maxArea = max(maxArea, area)
		}

	}
	// one by one pop remaining indices in stack
	// stop when only -1 is left (cannot use isEmpty)
	for len(st) > 1 {
		tpIdx, _ := st.Pop()
		currHt := heights[tpIdx]
		// note: index = len(heights) here
		area := currHt * (index - st.Peek() - 1)
		maxArea = max(maxArea, area)
	}

	return maxArea
}

// tricky solution: time = O(n) space = O(n)
// for each bar calculate area assuming it's the min bar
// for that you have to remember the last seen min and the next min bar
// if you have these two min indices on each side, area with current bar
// curr_area = bar_ht * (right_min - left_min - 1)
// to solve this is linear time we need a stack
// we push the bar indices sees in increasing order
// when we see a lower  bar (decreasing sequence), we pop the stack
// till we see a lower bar index
// when we pop each bar we calculate the area based on the left min and right min
// when all elements are pushed, we pop the stack one by one till empty following the above steps
// at the end we keep track of the max area seen so far
