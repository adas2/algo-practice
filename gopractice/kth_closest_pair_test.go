package practice

import (
	"fmt"
	"testing"
)

func TestSmallestDistancePair(t *testing.T) {
	nums := []int{1, 6, 1}
	k := 3
	// nums2 := []int{4, 1, 2, 2}
	fmt.Println(smallestDistancePair(nums, k))
}
