package dp

import (
	"fmt"
	"testing"
)

func TestLengthofLIS(t *testing.T) {
	// Input: [10,9,2,5,3,7,101,18]
	// Output: 4
	inp := []int{10, 9, 2, 5, 3, 7, 101, 18}

	fmt.Println("LIS:", lengthOfLIS(inp))
}

// patience sorting
// 2,3,7,18

func TestEfficientLIS(t *testing.T) {
	inp := []int{10, 9, 2, 5, 3, 7, 101, 18}

	fmt.Println("LIS:", efficientLIS(inp))
}
