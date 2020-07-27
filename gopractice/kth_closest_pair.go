package practice

import (
	"fmt"
	"sort"
)

// LC # 719 (hard)

func smallestDistancePair(nums []int, k int) int {

	n := len(nums)
	if k > n*(n-1)/2 {
		// invalid k if > Nc2
		fmt.Printf("invalid k\n")
		return -1
	}
	// sort the array and get the max pair-dist
	sort.Ints(nums)

	// now make a guess in this range and find out num pairs with dist <= guess
	hi := nums[n-1] - nums[0]
	lo := 0
	mid, count := 0, 0

	// binary search the guess
	for lo < hi {
		mid = (hi + lo) / 2
		fmt.Printf("hi, lo, guess: %d %d %d\n", hi, lo, mid)
		count = checkGuess(nums, mid)
		fmt.Printf("count = %d\n", count)
		if count >= k {
			// too high guess, reduce the guess
			// candidate = mid
			hi = mid
		} else {
			// to low increase the guess
			lo = mid
			// candidate = mid + 1
		}
	}

	// note lo is return, because either lo=mid when count > k or lo=mid+1 when count < k
	return lo
}

// helper to return the number of pairs with dist <= guess
func checkGuess(nums []int, guess int) int {
	count, i := 0, 0
	// use 2 ptr-window mechanism
	// note although 2 for loops its one traversal so O(n)
	for j := range nums {
		// if the range ti too large move left pointer
		for (nums[j] - nums[i]) > guess {
			i++
		}
		// count increases by farthest index reached (tricky!)
		count += j - i
	}
	return count
}

// Given the int array (unsorted), find the closest pair
// Strategy: Naive solution is O(n^2) time and space
// To improve use trail and error method, i.e. guess a dist value
// check if it is too much or too less
// to get the pair distances, use sorting anf 2 ptr window
// E.g. sorted arr := {1,2,2,4,7}
// max dist = 6(7-1); start with i,j=0, incr j till dist exceeds expected dist
// if dist > guess, move i to right so that it is decremented
// IMO: this is the most tricky part of the problem
// Rest is similar to binary search
