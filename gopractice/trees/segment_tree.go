package trees

import (
	"fmt"
	"math"

	"adas2.io/practice/dp"
)

// Description: Max segment tree definition

// define segment tree type
type segementTree []int

// iniaize the struct
func initSegmentTree(nums []int) segementTree {
	// create the empty data array with len 2n
	n := len(nums)
	// default inits to zero
	segTree := make(segementTree, n)
	// note node/index 0 is unitilized i.e dont care
	segTree = append(segTree, nums[:]...)
	fmt.Printf("Seg Tree: %v\n", segTree)
	// update the elements 1 to n-1 with the max value
	for i := n - 1; i > 0; i-- {
		segTree[i] = dp.Max(segTree[2*i], segTree[2*i+1])
	}

	// return
	return segTree
}

// updates a value in array[x] with v
func (s segementTree) update(nums []int, x, v int) segementTree {
	// change the value at respective index in seg tree
	index := x + len(nums)
	s[index] = v
	// percolate the changes to parent nodes
	// avoid node 0
	for index > 1 {
		index /= 2
		s[index] = dp.Max(s[2*index], s[2*index+1])
	}
	// updated tree
	return s
}

// range query max over a given range [x,y] inclusive
func (s segementTree) query(nums []int, x, y int) int {
	// optimzation: when whole len of array, output element segTree[1]
	if y-x+1 == len(nums) {
		return s[1] //contains entire range
	}
	// for all other case traverse the tree bottom up: O(logN) time
	left, right := len(nums)+x, len(nums)+y
	maxNum := math.MinInt32

	for left <= right {
		fmt.Printf("left: %v, right: %v\n", left, right)
		// if left index is odd, out of range, calc max and index++
		if left%2 == 1 {
			maxNum = dp.Max(maxNum, s[left])
			left++
		}
		// if right is even, i.e out of range, calc max and index--
		if right%2 == 0 {
			maxNum = dp.Max(maxNum, s[right])
			right--
		}

		fmt.Println("maxnum:", maxNum)
		// move one level up in seg tree
		left /= 2
		right /= 2
	}

	return maxNum
}
