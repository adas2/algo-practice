package practice

// Given a set of N (sorted) coordinates of stables and k cows find the maximum separation
// that can be achieved by after placing all cows among those coordinates

// FindMaxSeparation finds max sep range possible in array with k occupants
func FindMaxSeparation(coords []int, k int) int {
	maxsep := coords[len(coords)-1] - coords[0]
	var candidate int
	// max possible separation if first and low cows are placed at start and end
	// remaining k-2 will be placed somehwere in between
	if k > 2 {
		maxsep = maxsep / (k - 1)
	} else { // 2 or less whole range is max sep
		return maxsep
	}
	// search values from maxsep to minsep (min possinle 1 for int arr)
	for candidate = maxsep; candidate >= 1; candidate-- {
		// check if this separation value has a feasible placement
		if IsFeasible(coords, candidate, k) {
			break
		}
	}

	return candidate
}

// IsFeasible checks is allocation with sep separation is feasible
func IsFeasible(arr []int, sep int, num int) bool {

	// previous postions to compare with
	prevIdx := 0
	index := -1
	var maxSepAchieved int = arr[len(arr)-1] - arr[0] //default

	// iterate on number of cows left (first 0 and last k-1 already placed)
	for i := 1; i < num-1; i++ {
		// using binary search find closest coordinates that is >= target
		index = BinarySearchLast(arr, arr[0]+(i*sep), index+1, len(arr)-1)
		if index == -1 {
			return false
		}
		// find the separation achieved
		maxSepAchieved = Min(arr[index]-arr[prevIdx], maxSepAchieved)

		prevIdx = index

	}
	// last element separation
	maxSepAchieved = Min(arr[len(arr)-1]-arr[index], maxSepAchieved)

	if maxSepAchieved < sep {
		return false
	}
	// if all intervals passed separation test
	return true
}

// Min is util func
func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
