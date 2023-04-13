package practice

import "fmt"

// LC # 324 wiggle-sort II
// Note wiggle sort I is impl in C++

// TODO: fix this soln
func wiggleSort(nums []int) {

	// Using quickselect find the median
	n := len(nums)
	// median value is the (n+1)/2 element in sorted order
	median := findKthSmallestValue(nums, (n+1)/2, 0, n-1)

	fmt.Println(median)

	// curr := 0

	var smallIndex int
	if n%2 == 1 {
		smallIndex = n - 1
	} else {
		smallIndex = n - 2
	}
	largeIndex := 1

	for curr := 0; curr < n; curr++ {

		// (1) elements smaller than the 'median' are put into the last even slots
		if nums[curr] < median && curr%2 != 0 {

			if (curr >= smallIndex) && (curr%2 == 0) {
				// skip already covered indices
				continue
			}
			// else swap smallIndex, curr
			nums[curr], nums[smallIndex] = nums[smallIndex], nums[curr]
			smallIndex -= 2
			curr--

		} else if nums[curr] > median {
			// (2) elements larger than the 'median' are put into the first odd slotsZ

			if (curr <= largeIndex) && (curr%2 != 0) {
				// skip laready covered indices
				continue
			}
			// swap laregIndx, curr
			nums[curr], nums[largeIndex] = nums[largeIndex], nums[curr]
			largeIndex += 2
			curr--

		}
	}

}

func findKthSmallestValue(nums []int, k, start, end int) int {
	// base case
	if start == end {
		return nums[start]
	}

	// parition using random quiselect
	index, val := quickSelect(&nums, start, end)

	// count of elements in left side of parition
	count := (index - start) + 1

	if count == k {
		return val
	}

	if k < count {
		return findKthSmallestValue(nums, k, start, index-1)
	} else {
		// if k > count
		return findKthSmallestValue(nums, k-count, index+1, end)
	}

}

// Very tricky, chosen from LC upvoted solution
// note elements are filled from the beginning and the end simultaneously, both end and both begienning strategies don't work
