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
	minDiff := math.Inf(1)
	sum, closestSum := 0, 0

	// Optional: keep track of the optimal indices
	indices := make([]int, 3)

	// for each index use pivot
	for i := range nums {
		// fix i and search 2 indices(j, k) such that nums[i]+ nums[j]+nums[k] ~ target
		nTarget := target - nums[i]
		j, k := i+1, len(nums)-1
		// make sure i is not counted twice
		for j < k {
			sum = nums[j] + nums[k] + nums[i]
			diff := math.Abs(float64(target - sum))
			if diff < minDiff {
				minDiff = diff
				closestSum = sum
				indices = []int{nums[i], nums[j], nums[k]}
				fmt.Printf("saw a new min: %v sum: %d\n", minDiff, sum)
			}

			if nTarget > nums[j]+nums[k] {
				j++
			} else {
				k--
			}

		}

	}
	fmt.Println("Final set of nums: ", indices)
	return closestSum
}

// algo:
// use the optimal 2sum approach: i.e.,
// sort the array and them use 2 pointers to calculate duplex sum
// to incorporate triplet, do the above for all elements
// time complexity: sort -> O(nlogn) + compare for n nums-> O(n*n)
// i.e., T = O(n^2), S=O(1)

// Regular 3sum
// LC #15
// Note: Fix: no duplicates to be printed
func threeSum(nums []int) [][]int {
	// sort the array
	sort.Ints(nums)
	// minDiff := math.Inf(1)
	sum := 0

	// set impl with triplet as key
	indexMap := make(map[indices]struct{})

	// for each index use pivot
	for i := range nums {
		nTarget := -nums[i]
		j, k := i+1, len(nums)-1
		for j < k {
			sum = nums[j] + nums[k] + nums[i]
			if sum == 0 {
				indexMap[indices{nums[i], nums[j], nums[k]}] = struct{}{}
			}

			if nTarget > nums[j]+nums[k] {
				j++
			} else {
				k--
			}

		}
	}

	triplets := make([][]int, 0)
	for key := range indexMap {
		triplets = append(triplets, []int{key.i, key.j, key.k})
	}
	return triplets
}

type indices struct {
	i, j, k int
}

// Note: this could be optimized further for space if we skip ahead repeated elements in original array
// that way we do not need a map traversal and we avpid duplicates while reading the array
