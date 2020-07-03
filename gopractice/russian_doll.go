package practice

import (
	"fmt"
	"sort"
)

// LC #354 russian doll

// returns the num of max ecvelopes
func maxEnvelopes(envelopes [][]int) int {
	// sort by width first
	sort.SliceStable(envelopes, func(i, j int) bool { return envelopes[i][0] < envelopes[j][0] })
	fmt.Println("Sorted by weights:", envelopes)

	// reverse sort for same widths
	sort.SliceStable(envelopes, func(i, j int) bool {
		return envelopes[i][0] == envelopes[j][0] && envelopes[i][1] > envelopes[j][1]
	})
	fmt.Println("Reverse Sorted by heights:", envelopes)

	// now calculate LIS on the heights
	res := findMaxLIS(envelopes)

	return res
}

// patience sort version
func findMaxLIS(envs [][]int) int {
	// buckets to keep decresing seqs
	buckets := []int{}

	// keep inserting/replacing element to sorted buckets
	for i := range envs {
		index := sort.SearchInts(buckets, envs[i][1])
		if index < len(buckets) {
			// match, replace
			buckets[index] = envs[i][1]
		} else {
			// add
			buckets = append(buckets, envs[i][1])
			// track the env seq
			if index > 0 {
				fmt.Println("last seq height:", buckets[index-1])
			}
		}
	}

	return len(buckets)
}

// algorithm:
//
// there a two dimensions height and weight ()
// first we sort against one dimensions
// e.g. (w1,h1) (w2,h2) (w3,h3)....
// sort by weight : [[5,4],[6,4],[6,7],[2,3]]
// [2,3] [5,4] [6,4] [6,7]
// then find the longest increasing subsequence of heights keeping sorted order
// special  case: for entries with same height e.g. [[2,2] [3,3] [3,4] [3,5] [4,5]]
// this will produce incorrect subsequence len like 2,3,4,5 --> 4 envelops
// this is because 3,4,5 all with same width adds to the subsequence
// to avoid this one intution is to reverse sort the heights for the same width
// this will force the LIS to choose the appropriate height
// e.g. [[2,2] [3,5] [3,4] [3,3] [4,5]]
// LIS: 2,4,5
// complexity: sort: O(nlogn) + LIS: O(n^2) or O(nlogn) with patient sort
