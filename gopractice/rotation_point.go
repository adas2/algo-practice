package practice

// Problem:  Given a list of strings on lenght N, where the first i elements are sorted
//  and next (N-i) are sorted independently, and elemnent[N-1] < element[0]
//  The rotation point of the array can be defined as the index of the lowest value
//	element. Find an algorithm to find the the rotation point of the given array.
//  Assume N can be very large, and no strings are repeated.
//  E.g. strList = ["cat", "elephant", "horse", "apple", "bamboo", "banana"]
//  Output = 3 i.e. rotation point

func findRotationPoint(strList []string) int {

	// size
	n := len(strList)

	// error case
	if n == 0 {
		return -1
	}

	var low, high, mid int
	low = 0
	high = n - 1

	// iterative binary search
	candidate := 0

	// if whole array is sorted
	if strList[0] < strList[n-1] {
		return candidate
	}

	for low <= high {
		mid = (low + high) / 2
		// check which half has the rotation point
		if strList[low] > strList[mid] {
			// iterate on left half
			candidate = high
			high = mid
		} else if strList[mid+1] > strList[high] {
			// iterate on right half
			candidate = high
			low = mid + 1
		} else {
			// this case will arise when either following conditions arise:
			// A. the mid+1 is rotation point,
			if low != high {
				candidate = mid + 1
				break
			}
			// B. low==high,  handled by other conditions

		}
	}

	// return the candidate
	return candidate

}
