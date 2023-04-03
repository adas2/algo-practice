package practice

import (
	"math"
	"sort"
)

func increasingTriplet(nums []int) bool {

	// monitor the smallest and second smallers int
	// if any integer is > second smallest return true

	smallest, secondSmallest := math.MaxInt32, math.MaxInt32

	for _, v := range nums {
		if v <= smallest {
			smallest = v
		}
		if v > smallest && v < secondSmallest {
			secondSmallest = v
		}

		if v > secondSmallest {
			return true
		}
	}

	return false
}

// Another approach using patience sorting (Note: this is similar to LIS problem)
// search op O(logn) with O(n) space of arr (here n = 3 elements)
// since at max arr will have 3 elements in arr, conplexity become O(n), and space is O(1)
func increasingTriplet2(nums []int) bool {
	arr := make([]int, 0)

	for i := range nums {
		// first element
		if i == 0 {
			arr = append(arr, nums[i])
			continue
		}

		// if the curr num is > last arr index
		// find the index to insert in the sorted order
		index := sort.SearchInts(arr, nums[i])
		if index <= len(arr)-1 {
			arr[index] = nums[i]
		} else {
			arr = append(arr, nums[i])
		}

		// if the arr len is >= 3 triplet exists
		if len(arr) >= 3 {
			return true
		}

	}

	return false

}
