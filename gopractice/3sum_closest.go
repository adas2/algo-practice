package practice

import (
	"fmt"
	"math"
	"sort"
)

// closest 3sum problem:
// find triplet such target - (nums[i]+nums[j]+nums[k]) is min
func threeSumClosest(nums []int, target int) int {
	// sort the array
	sort.Ints(nums)
	// fmt.Println(nums)
	minDiff := math.MaxInt32
	sum, closestSum := 0, 0
	// for each num use pivot
	for i := range nums {
		// fix i and search 2 indices(j, k) such that nums[i]+ nums[j]+nums[k] ~ target
		nTarget := target - nums[i]
		j, k := i+1, len(nums)-1
		// make sure i is not counted twice
		for j < k {
			// fmt.Printf("i: %d, j: %d, k: %d, target: %d\n", i, j, k, nTarget)
			sum = nums[j] + nums[k] + nums[i]
			diff := abs(target - sum)
			if diff < minDiff {
				minDiff = diff
				closestSum = sum
				fmt.Printf("saw a new min: %d sum: %d\n", minDiff, sum)
			}

			if nTarget > nums[j]+nums[k] {
				j++
			} else {
				k--
			}

		}

	}
	return closestSum
}

// util
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// algo:
// use the optimal 2sum approach: i.e.,
// sort the array and them use 2 pointers to calculate duplex sum
// to incorporate triplet, do the above for all elements
// time complexity: sort: O(nlogn) + compare for n nums: O(n*n) = O(n^2) total
// Abs returns the absolute value of x.
