package dp

import (
	"fmt"
	"sort"
)

// find the longest increasing subsequence in a given array (int)

// method 1: usng dp with T=O(N^2), S=O(N)
// output: length of the LIS
func lengthOfLIS(nums []int) int {
	// nil case
	if len(nums) == 0 {
		return 0
	}

	// init dp array
	dp := make([]int, len(nums))
	// len 1 lis = 1
	dp[0] = 1
	// max lis seen so far (default 1)
	max_seen := 1

	for i := 1; i < len(nums); i++ {
		// case 1: ith num is included in lis (provided condition satisfied)
		// condition: there exists at least nums[j], j<i, st nums[j] < nums[i]
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = Max(dp[i], dp[j]+1)
			}
		}
		// case 2: ith num starts a new lis
		dp[i] = Max(dp[i], 1)
		max_seen = Max(max_seen, dp[i])
	}

	return max_seen
}

// explanation (method 1):
// dp[i] ==> LIS len for arr[0..i] with sequence ending in i
// dp[i] ==> Max(DP[j])+1, for all 0<=j<i and arr[i]>arr[j])
// else if no larger element dp[i]=dp[i-1]

// method 2: using patience sorting
// T=O(nlogn) S=O(n)
func efficientLIS(nums []int) int {
	// create empty sequence array
	seq := make([]int, 0)

	for _, v := range nums {

		// add num if it is first or larger than last num
		// note seq is always sorted
		if len(seq) == 0 || seq[len(seq)-1] < v {
			seq = append(seq, v)
			// Optional: record the element which caused a new entry
			// this helps to get an LIS
			if len(seq) > 1 {
				fmt.Printf("seq num: %d\n", seq[len(seq)-2])
			}
		} else {
			// find the right index using bisnary search
			index := sort.SearchInts(seq, v)
			seq[index] = v
		}

	}
	// final number in LIS
	fmt.Printf("seq num: %d\n", seq[len(seq)-1])

	return len(seq)

}

// InsertSorted inserts into sorted array maintaining its property
func InsertSorted(s []int, e int) []int {
	s = append(s, 0)
	i := sort.Search(len(s), func(i int) bool { return s[i] > e })
	copy(s[i+1:], s[i:])
	s[i] = e
	return s
}

// explanation (method 2):
// use patience sorting to stack the elements in the order they appear
// while there are no numbers left do the following:
// 1. place the first number on a new stack
// 2. one by one place the numbers on the leftmost stack whose top > its value
// 3. if not such stack create a new one
// 4. when no nums left calculate the num of stacks created
// to get a LIS sequence, backtrack using the stack vals

// method 3: using segment trees
// similar in complexity different implementation
// find relative index in the array
// use count segment tree
