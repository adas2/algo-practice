package practice

import (
	"fmt"
	"sort"
)

// LC # 56 (merge overlapping intervals)

// simple impl with sorted slices ans using two stacks
// intervals[i] = [starti, endi]
// T = O(NlogN) + O(N)
// S = O(N)
func merge(intervals [][]int) [][]int {

	// create a sorted list of events with <start_time, event_id_id, start/end>
	s := []ts{}

	// populate list
	for i, ev := range intervals {
		s = append(s, ts{ev[0], i, "start"})
		s = append(s, ts{ev[1], i, "end"})
	}
	// process the above in increasing order of start_times
	sort.SliceStable(s, func(i, j int) bool {
		if s[i].time == s[j].time {
			return s[i].flag > s[j].flag // "start" > "end"
		}
		return s[i].time < s[j].time
	})
	// tricky: to handle events that are adjacent (start_i+1 == end_i), customize such that, start_i+1 comes first
	fmt.Println(s)

	// maintain set to keep track of overlapping events
	eSet := make(map[int]struct{})

	// populate result array
	res := [][]int{}
	c_start, c_end := 0, 0

	for _, e := range s {

		// event exists in set create or delet based on start or end
		if _, exists := eSet[e.event_id]; !exists && e.flag == "start" {
			if len(eSet) == 0 {
				c_start = e.time
			}
			eSet[e.event_id] = struct{}{}
		} else if e.flag == "end" {
			c_end = e.time
			delete(eSet, e.event_id)
		}
		// else case: error case

		// if set is empty no overlapping event is complete
		if len(eSet) == 0 {
			res = append(res, []int{c_start, c_end})
		}

	}

	return res
}

// timestap
type ts struct {
	time     int
	event_id int
	flag     string
}

// Optimized implementation
// beautiful and elegent :)
func merge2(intervals [][]int) [][]int {

	// sort like previous one
	sort.SliceStable(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	fmt.Println(intervals)

	res := [][]int{}
	// traverse and merge the intervals based on the following:
	for _, interval := range intervals {
		// if no entry in the res, just add the interval
		// or if end time of the last merged interval in res < start time of the current interval
		if len(res) == 0 || res[len(res)-1][1] < interval[0] {
			res = append(res, interval)
		} else {
			// there is overlap between last interval in res and current interval
			// merge such that end time of last entry is changed to max(last, current)
			res[len(res)-1][1] = max(res[len(res)-1][1], interval[1])
		}
	}

	return res
}
