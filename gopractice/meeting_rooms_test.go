package practice

import (
	"fmt"
	"testing"
)

// [[0,30],[5,10],[15,20]]
func TestMinMeetingRooms(t *testing.T) {
	intervals := [][]int{
		{0, 30},
		{5, 10},
		{15, 20},
	}

	fmt.Println(minMeetingRooms(intervals))
}
