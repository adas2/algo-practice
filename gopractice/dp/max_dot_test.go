package dp

import (
	"fmt"
	"testing"
)

func TestMaxDotProduct(t *testing.T) {
	// [-3,-8,3,-10,1,3,9]
	// [9,2,3,7,-9,1,-8,5,-1,-1]

	nums1 := []int{-3, -8, 3, -10, 1, 3, 9}
	nums2 := []int{9, 2, 3, 7, -9, 1, -8, 5, -1, -1}

	result := maxDotProduct(nums1, nums2)
	fmt.Println(result)
	// assert.NotNil(t, result)
}
