package practice

import (
	"fmt"
	"testing"
)

func TestFindKthLargest(t *testing.T) {
	nums := []int{3, 2, 1, 6, 5}
	k := 2

	fmt.Printf("Nums: %v\nK: %d, K-th largest value: %d\n", nums, k, findKthLargest(nums, k))
}
