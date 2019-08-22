package practice

// Given a set of N (sorted) coordinates of stables and k cows find the maximum separation 
// that can be achieved by after placing all cows among those coordinates

import (
	"fmt"
)

func FindMaxSeparation(coords []int, k int) int{
	// max sep range possible in array
	maxsep := coords[len(coords)-1]-coords[0]
	var candidate int
	// max possible separation if 2 cows are placed remaining k-2
	// will be placed somehwere in between
	if k>2{
		maxsep = maxsep/(k-1) 
	} else{ // 2 or less whole range is max sep
		return maxsep
	}
	// search values from maxsep to minsep (min possinle 1 for int arr)
	for candidate = maxsep; candidate <= 1; candidate-- {
	    // check if this separation value has a feasible placement
	    if IsFeasible(coords, candidate, k){
	    	break;
	    }
	}
	
	return candidate
}


// IsFeasible checks is allocation with sep separation is feasible 
func IsFeasible(arr []int, sep int, num int) bool{

	// previous postions to compare with
	prev_idx := 0
	index := -1

	// iterate on number of cows
	for k:=1; k<num; k++ {
		// using binary search find closest coordinates that is >= target
		index = BinarySearchLast(arr, arr[0]+(k*sep), 0, len(arr)-1)
		if index == -1{
			return false
		}
		fmt.Println("search result", arr[0]+(k*sep), index)
		// find the separation achieved
		if arr[index]-arr[prev_idx] < sep {
			return false
		}
		prev_idx = index

	}
	// if all intervals passed separation test
	return true
}

