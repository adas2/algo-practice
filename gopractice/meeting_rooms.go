package practice

import (
	"fmt"
	"sort"
)

// O(n^2longn) solution better soln below
func minMeetingRooms(intervals [][]int) int {
	// sort using start times
	sort.SliceStable(intervals, func(i, j int) bool {
		// for same start times, sort by increasing end times
		// if intervals[i][0] == intervals[j][0] {
		// 	return intervals[i][1] > intervals[j][1]
		// }
		return intervals[i][0] < intervals[j][0]
	})

	// walk through the array and find the conflict and add extra room
	// rooms := 1
	roomSet := []int{}

	// for i := 0; i < len(intervals)-1; i++ {
	for _, interval := range intervals {

		// find the lowest index in room set
		index := sort.SearchInts(roomSet, interval[0])

		fmt.Printf("Start time: %d,  Search index: %d\n", interval[0], index)
		// TODO: this needs to be changed to priority queue, since we need to find the element with the earliest finish time
		// it the current start time > earliest finish time, then we can reuse a meeting room, else we need to create a new room
		if index < len(roomSet) && roomSet[index] <= interval[0] {
			// found an interval which ends before start
			roomSet[index] = interval[1]
		} else if index != 0 {
			roomSet[index-1] = interval[1]
		} else {
			// index = len(roomSet)
			roomSet = append(roomSet, interval[1])
		}

		sort.SliceStable(roomSet, func(i, j int) bool { return roomSet[i] < roomSet[j] })

		// if no entry in rooms or a conflict with next event add a room
		// if intervals[i+1][0] < intervals[i][1] {
		// 	rooms++
		// }
		fmt.Println(roomSet)

	}
	return len(roomSet)
}

// logic: using a stack set:
// sort in start times, place the first slot on the left most stack bucket
// pick the leftmost index i from stack such value[i] <= curr_start_time,
// update the value[i] to curr_end_time
// return num buckets
// Example: [[0,30],[5,10],[[0,30],[5,10],[15,70]]
// roomset[0] = 30, search(5) --> -1, i.e. < len(roomset), insert_sorted
// roomset[0]-->10, roomset[1]-->30
// search(15) --> 10, i.e. 0, update roomset[0] = 70
// return len(roomset)
/**
class Solution {
public:
int minMeetingRooms(vector& intervals) {
map<int, int> mp; // key: time; val: +1 if start, -1 if end

    for(int i=0; i< intervals.size(); i++) {
        mp[intervals[i].start] ++;
        mp[intervals[i].end] --;
    }

    int cnt = 0, maxCnt = 0;
    for(auto it = mp.begin(); it != mp.end(); it++) {
        cnt += it->second;
        maxCnt = max( cnt, maxCnt);
    }

    return maxCnt;
}
};
**/

// nice ane elegant O(nlogn) solution
func minMeetingRooms2(intervals [][]int) int {

	// create a map to track all start and end times for each intervals
	iMap := make(map[int]int)

	// for all intervals
	for _, interval := range intervals {
		iMap[interval[0]] += 1
		iMap[interval[1]] -= 1
	}

	// sort the times in map in ascending order
	keys := make([]int, 0)
	for key := range iMap {
		keys = append(keys, key)
	}
	sort.SliceStable(keys, func(i, j int) bool { return keys[i] < keys[j] })

	cnt, maxCnt := 0, 0
	for _, k := range keys {
		cnt += iMap[k]
		maxCnt = max(cnt, maxCnt)
	}

	return maxCnt
}
