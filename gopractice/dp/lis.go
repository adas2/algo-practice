package dp

import (
	"fmt"
	"sort"
)

// find the longest increasing subsequence in a given array (int)

// method 1: usng dp with O(N^2) complexity
// output: length of the LIS
func lengthOfLIS(nums []int) int {
	// error case
	if len(nums) == 0 {
		return 0
	}
	// adjust size to include whole in array
	dp := make([]int, len(nums))
	// base case: LIS len for only 1 element
	dp[0] = 1
	maxLIS := 0 //init
	result := 1

	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				// new sequence foound
				maxLIS = Max(maxLIS, dp[j])
			}
		}
		// in case no sequence, LIS is 1
		dp[i] = Max(maxLIS+1, 1)
		// fmt.Printf("maxLIS: %d, dp[%d] = %d\n", maxLIS, i, dp[i])
		// store the max len seen so far
		result = Max(result, dp[i])
		// maxLIS = Max(maxLIS+1, dp[i-1])
	}

	return result
}

// explanation:
// dp[i] ==> LIS len for arr[0..i] with sequence ending in i
// dp[i] ==> Max(DP[j])+1, for all 0<=j<i and arr[i]>arr[j])
// else if no larger element dp[i]=dp[i-1]

// method 2: using patience sorting
// O(n log n)
func efficientLIS(nums []int) int {
	// create empty sequence array
	seq := make([]int, 0)

	for i, v := range nums {
		// first num
		if i == 0 {
			seq = append(seq, v)
			continue
		}
		// fmt.Println(seq)
		// find the closest index < v
		index := sort.SearchInts(seq, v)
		if index < len(seq) {
			seq[index] = v
		} else {
			// insert the num
			seq = append(seq, v)
			// record the element which caused a new entry
			fmt.Printf("seq num: %d\n", seq[len(seq)-2])
		}
		// seq = InsertSorted(seq, nums[i])
		fmt.Printf("index: %d, sequence array: %v\n", index, seq)

	}

	// final seq number
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

// explanation:
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
