package practice

import (
	"fmt"
	"testing"
)

func TestMergeIntervals(t *testing.T) {
	// [[1,3],[2,6],[8,10],[15,18]]
	ints := [][]int{{8, 10}, {1, 3}, {2, 6}, {10, 15}}

	fmt.Println(merge2(ints))
}
