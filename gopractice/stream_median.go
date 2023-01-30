package practice

import custom "adas2.io/practice/custom_impl"

// find the median from input stream of data

// MedianFinder is soln
type MedianFinder struct {
	lower  *custom.MaxHeap
	higher *custom.MinHeap
}

// Constructor initialize your data structure here.
func Constructor() MedianFinder {
	// declare empty arrays for both halves
	l, h := []int{}, []int{}
	minH := custom.InitMinHeap(h)
	maxH := custom.InitMaxHeap(l)
	// create struct
	return MedianFinder{
		lower:  maxH,
		higher: minH,
	}
}

// AddNum adds
func (m *MedianFinder) AddNum(num int) {
	// add such that balance is maintained between lower and higher
	// i.e. either len(lower) == len(higher) or len(lower) = len(higher) + 1 (odd case)
	m.lower.Add(num)

	// balance
	if m.lower.Len() > m.higher.Len()-1 {
		v := m.lower.GetMax()
		m.higher.Add(v)
	}

	// if higher becomes larger
	if m.higher.Len() > m.lower.Len() {
		v := m.higher.GetMin()
		m.lower.Add(v)
	}
}

// FindMedian finds
func (m *MedianFinder) FindMedian() float64 {
	// even case: take max of lower and min of
	if (m.lower.Len()+m.higher.Len())%2 == 0 {
		return float64(m.lower.Peek()+m.higher.Peek()) / 2
	}
	// odd case: max of lower
	return float64(m.lower.Peek())
}
