package strs

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// Rearrange a list of numbers to get the largest number by concating
// LC # 179

// custom sort based soln based on lc editorial
func largestNumber(nums []int) string {

	// convert to str
	numStrs := make([]string, len(nums))
	for i := range nums {
		numStrs[i] = strconv.Itoa(nums[i])
	}
	fmt.Printf("nums_strs:  %v\n", numStrs)

	cmp := func(i, j int) bool {
		order1 := numStrs[i] + numStrs[j]
		order2 := numStrs[j] + numStrs[i]
		return order1 > order2
	}

	// descending sort based on custom comparator
	sort.SliceStable(numStrs, cmp)

	res := strings.Join(numStrs, "")

	// to handle the case all zeros e.g. [0,0,0]
	if res[0]-'0' == 0 {
		return "0"
	}

	return res
}
